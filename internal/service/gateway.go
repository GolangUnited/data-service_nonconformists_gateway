package service

import (
	"errors"
	"fmt"
	"go-united/gateway/internal/structs"
	"log"
	"time"

	"github.com/RealJimy/gounited-api/grpc/certificates"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type Gateway struct {
	cliCfg             map[string]structs.ClientConfig
	connections        map[string]*grpc.ClientConn
	ClientCertificates certificates.CertificatesClient
}

const (
	serv_certificates string = "certificates"
	serv_courses      string = "courses"
)

var kacp = keepalive.ClientParameters{
	Time:                30 * time.Second, // send pings every 30 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func Init(cfg map[string]structs.ClientConfig) Gateway {
	gw := Gateway{
		cliCfg: cfg,
	}
	gw.connectAllClients()
	return gw
}

func (gw *Gateway) GetCertificates(userId string) ([]structs.Certificate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := gw.ClientCertificates.List(ctx, &certificates.ListRequest{UserId: &userId})
	if err != nil {
		log.Printf("certificate list error: %v", err)
		return []structs.Certificate{}, errors.New("internal error")
	}

	if len(resp.Certificates) == 0 {
		return []structs.Certificate{}, nil
	}

	res := make([]structs.Certificate, len(resp.Certificates))
	for i, cert := range resp.Certificates {
		res[i] = structs.Certificate{
			Id:        cert.Id,
			CreatedAt: uint32(cert.CreatedAt.Seconds),
		}
	}
	return res, nil
}

func (gw *Gateway) DisconnectAllClients() {
	for _, conn := range gw.connections {
		if conn != nil {
			conn.Close()
		}
	}
}

func (gw *Gateway) connectAllClients() {
	gw.connections = make(map[string]*grpc.ClientConn)
	gw.connections[serv_certificates] = gw.connectBackend(serv_certificates)
	gw.ClientCertificates = certificates.NewCertificatesClient(gw.connections[serv_certificates])
}

func (gw *Gateway) connectBackend(serviceName string) *grpc.ClientConn {
	connectionString := gw.prepareConnString(serviceName)
	connection, err := grpc.Dial(
		connectionString,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp),
	)
	if err != nil {
		log.Printf("error while connecting service `%s`: %v", connectionString, err)
		return nil
	}
	return connection
}

func (gw Gateway) prepareConnString(name string) string {
	connParams := gw.cliCfg[name]
	return fmt.Sprintf("%s:%s", connParams.Host, connParams.Port)
}
