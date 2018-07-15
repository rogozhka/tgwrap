package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Token  string
	ChatID string
	FromChatID string
	MessageID uint64
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
	FromChatID = viper.GetString("FromChatID")
	MessageID = uint64(viper.GetInt64("MessageID"))
}
