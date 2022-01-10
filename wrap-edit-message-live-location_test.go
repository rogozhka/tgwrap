package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EditMessageLiveLocationWrap(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)
	opt := SendLocationOpt{
		LivePeriod: 86000,
	}
	res, err := bot.SendLocation(chatID, 41.89, 12.50, &opt)
	is.Nil(err, "SendLocation err")

	longitude := 12.45
	latitude := 41.88

	_, err4 := bot.EditMessageLiveLocation(latitude, longitude,
		&EditMessageLiveLocationOpt{
			ChatID:    res.Chat.ID,
			MessageID: res.ID,
		})
	is.Nil(err4, "EditMessageLiveLocation err")
}
