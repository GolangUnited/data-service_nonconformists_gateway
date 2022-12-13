package vipercfg

import (
	"fmt"

	"github.com/spf13/viper"
)

type ViperCfg struct {
}

func New() ViperCfg {
	return ViperCfg{}
}

func (v ViperCfg) Fill(cfgData interface{}) {

	// TODO: Implement getting the config file name from flags
	// TODO: Implement reading config from env

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		// TODO: Separate config not found error
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Unmarshal(cfgData)
}
