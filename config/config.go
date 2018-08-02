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
	log.Println("initializing config using env variables")

	viper.SetEnvPrefix("tgwrap")

	viper.BindEnv("token")
	viper.BindEnv("ChatID")
	viper.BindEnv("FromChatID")
	viper.BindEnv("MessageID")

	Token = viper.GetString("Token")
	ChatID = viper.GetString("ChatID")
	FromChatID = viper.GetString("FromChatID")
	MessageID = uint64(viper.GetInt64("MessageID"))
}
