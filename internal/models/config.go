package models

type ServerConfig struct {
	RestPort string
	GrpcPort string
}

type CfgData struct {
	Server ServerConfig
}
