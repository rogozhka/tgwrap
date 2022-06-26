package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StopMessageLiveLocationWrap(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)
	opt := SendLocationOpt{
		LivePeriod: 86000,
	}
	res, err := bot.SendLocation(chatID, 41.89, 12.50, &opt)
	is.Nil(err, "SendLocation err")
	_, err2 := bot.StopMessageLiveLocation(
		&StopMessageLiveLocationOpt{
			ChatID:    res.Chat.ID,
			MessageID: res.ID,
		})
	is.Nil(err2, "StopMessageLiveLocation err")
}
