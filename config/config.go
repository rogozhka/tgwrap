package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Token  string
	ChatID string
)

func init() {

	log.Println("initializing config")

	viper.SetConfigFile("config/config.json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("config initiazlized using file: %s\n", viper.ConfigFileUsed())

	Token = viper.GetString("Token")
	ChatID = viper.GetString("ChatID")
}
