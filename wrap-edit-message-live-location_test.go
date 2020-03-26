package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EditMessageLiveLocationWrap(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	opt := SendLocationOpt{
		LivePeriod: 86000,
	}
	res, err := bot.SendLocation(chatID, 41.89, 12.50, &opt)
	assert.Nil(t, err, "SendLocation err")

	longitude := 12.45
	latitude := 41.88

	_, err4 := bot.EditMessageLiveLocation(latitude, longitude,
		&EditMessageLiveLocationOpt{
			ChatID:    res.Chat.ID,
			MessageID: res.ID,
		})

	assert.Nil(t, err4, "EditMessageLiveLocation err")
}
