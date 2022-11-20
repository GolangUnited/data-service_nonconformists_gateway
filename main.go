package main

import (
	"go-united/gateway/internal/api/rest"
	"go-united/gateway/internal/config"
	"go-united/gateway/internal/server"
	"go-united/gateway/pkg/config/vipercfg"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Init(vipercfg.New())
	echoInstance := echo.New()
	rest.RegisterHandlers(echoInstance, server.NewServer())
	echoInstance.Logger.Fatal(echoInstance.Start(":" + cfg.Server.RestPort))
}
