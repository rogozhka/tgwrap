package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendLocationWrap(t *testing.T) {
	token := getTokenEnv()
	bot := createBot(token)
	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	longitude := 12.50
	latitude := 41.89

	_, err4 := bot.SendLocation(chatID, latitude, longitude,
		&SendLocationOpt{
			DisableNotification: false,
			LivePeriod:          65,
		})

	assert.Nil(t, err4, "SendMessage err")
}
