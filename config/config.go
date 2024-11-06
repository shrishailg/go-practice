package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func Init() {
	Config = viper.New()
	Config.SetConfigType("json")
	Config.SetConfigName("app_config")
	Config.AddConfigPath("/config/")
	Config.AddConfigPath("./config/")
	Config.AddConfigPath("*/config/")

	if err := Config.ReadInConfig(); err != nil {
		log.Fatal("Error reading config", err.Error())
	}
}
