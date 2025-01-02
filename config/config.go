package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func LoadConfig() {
	Config = viper.New()
	Config.SetConfigName("app_config")
	Config.SetConfigType("json")
	Config.AddConfigPath("/config/")
	Config.AddConfigPath("./config/")
	Config.AddConfigPath("*/config/")

	if err := Config.ReadInConfig(); err != nil {
		log.Fatal("Error in intializing config", err.Error())
	}
}
