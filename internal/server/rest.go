package server

import (
	"fmt"
	"go-united/gateway/internal/api/rest"

	"github.com/labstack/echo/v4"
)

type RestServer struct {
	// serv *echo.Echo
}

func NewServer() *RestServer {
	return &RestServer{
		// serv: server,
	}
}

func (s RestServer) GetCertificates(ctx echo.Context, params rest.GetCertificatesParams) error {
	fmt.Println("GetCertificates not implemented")
	return nil
}

func (s RestServer) GetProfile(ctx echo.Context) error {
	fmt.Println("GetProfile not implemented")
	return nil
}
