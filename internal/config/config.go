package config

import (
	"go-united/gateway/internal/interfaces"
	"go-united/gateway/internal/structs"
)

func setDefaults(cfg *structs.CfgData) {
	cfg.Server.RestPort = "8081"
	cfg.Server.GrpcPort = "8082"
}

func Init(configer interfaces.ConfigI) structs.CfgData {
	cfg := structs.CfgData{}
	setDefaults(&cfg)
	configer.Fill(&cfg)
	return cfg
}
