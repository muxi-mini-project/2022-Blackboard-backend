package config

import (
	"log"

	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigFile("./conf/config.yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

}
