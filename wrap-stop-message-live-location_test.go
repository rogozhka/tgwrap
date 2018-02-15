package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StopMessageLiveLocationWrap(t *testing.T) {
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

	_, err4 := bot.StopMessageLiveLocation(
		&StopMessageLiveLocationOpt{
			ChatID:    res.Chat.ID,
			MessageID: uint(res.ID),
		})

	assert.Nil(t, err4, "StopMessageLiveLocation err")
}
