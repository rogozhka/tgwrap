package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendLocationWrap(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)

	longitude := 12.50
	latitude := 41.89

	_, e := bot.SendLocation(chatID, latitude, longitude,
		&SendLocationOpt{
			DisableNotification: false,
			LivePeriod:          65,
		})

	assert.Nil(t, e, "SendLocation err")
}

func Test_SendLocationWrap_outOfRange(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)

	longitude := 12.50
	latitude := 41.89

	_, err := bot.SendLocation(chatID, latitude, longitude,
		&SendLocationOpt{
			DisableNotification: false,
			LivePeriod:          10,
		})

	assert.NotNil(t, err, "SendLocation err")
}
