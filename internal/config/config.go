package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	viper.SetConfigName("env")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.WatchConfig()
	return viper.GetViper()
}
