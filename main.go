package main

import (
	"go-united/gateway/internal/api/rest"
	"go-united/gateway/internal/config"
	"go-united/gateway/internal/server"
	"go-united/gateway/internal/service"
	"go-united/gateway/pkg/config/vipercfg"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Init(vipercfg.New())

	echoInstance := echo.New()

	gateway := service.Init(cfg.Clients)
	defer gateway.DisconnectAllClients()

	restServer := server.New(gateway)

	rest.RegisterHandlers(echoInstance, restServer)

	echoInstance.Logger.Fatal(echoInstance.Start(":" + cfg.Server.RestPort))
}
