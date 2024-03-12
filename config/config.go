package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf *viper.Viper

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetDefault("server.port", 8080)

	if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Failed to read configuration file: %v", err)
    }

	Conf = viper.GetViper()
}