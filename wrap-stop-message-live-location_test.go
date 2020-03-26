package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StopMessageLiveLocationWrap(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	opt := SendLocationOpt{
		LivePeriod: 86000,
	}
	res, err := bot.SendLocation(chatID, 41.89, 12.50, &opt)
	assert.Nil(t, err, "SendLocation err")

	_, err2 := bot.StopMessageLiveLocation(
		&StopMessageLiveLocationOpt{
			ChatID:    res.Chat.ID,
			MessageID: res.ID,
		})

	assert.Nil(t, err2, "StopMessageLiveLocation err")
}
