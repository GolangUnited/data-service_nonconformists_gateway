package structs

type ServerConfig struct {
	RestPort string
	GrpcPort string
}

type ClientConfig struct {
	Host string
	Port string
}

type CfgData struct {
	Server  ServerConfig
	Clients map[string]ClientConfig
}
