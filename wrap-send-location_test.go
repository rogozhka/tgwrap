package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendLocationWrap(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)

	longitude := 12.50
	latitude := 41.89

	_, e := bot.SendLocation(chatID, latitude, longitude,
		&SendLocationOpt{
			DisableNotification: false,
			LivePeriod:          65,
		})

	is.Nil(e, "SendLocation err")
}

func Test_SendLocationWrap_outOfRange(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)

	longitude := 12.50
	latitude := 41.89

	_, err := bot.SendLocation(chatID, latitude, longitude,
		&SendLocationOpt{
			DisableNotification: false,
			LivePeriod:          10,
		})

	is.NotNil(err, "SendLocation err")
}
