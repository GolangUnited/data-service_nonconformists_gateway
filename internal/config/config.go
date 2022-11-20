package config

import (
	"go-united/gateway/internal/interfaces"
	"go-united/gateway/internal/models"
)

func setDefaults(cfg *models.CfgData) {
	cfg.Server.RestPort = "8081"
	cfg.Server.GrpcPort = "8082"
}

func Init(configer interfaces.ConfigI) models.CfgData {
	cfg := models.CfgData{}
	setDefaults(&cfg)
	configer.Fill(&cfg)
	return cfg
}
