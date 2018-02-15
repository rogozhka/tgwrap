package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EditMessageLiveLocationWrap(t *testing.T) {
	token, err := getTokenEnv()
	assert.Nil(t, err, "Token")
	if len(token) < 1 {
		return
	}
	bot := NewBot(token)

	arr, err2 := bot.GetUpdates(&GetUpdatesOpt{Limit: 1})
	assert.Nil(t, err2, "GetUpdates err")

	if len(arr) < 1 {
		return
	}

	up := arr[0]

	opt := SendLocationOpt{
		LivePeriod: 86000,
	}
	res, err := bot.SendLocation(up.Message.Chat.ID, 41.89, 12.50, &opt)

	longitude := 12.45
	latitude := 41.88

	_, err4 := bot.EditMessageLiveLocation(latitude, longitude,
		&EditMessageLiveLocationOpt{
			ChatID:    res.Chat.ID,
			MessageID: uint(res.ID),
		})

	assert.Nil(t, err4, "EditMessageLiveLocation err")
}
