package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUpdates(t *testing.T) {

	token := getTokenEnv()
	bot := createBot(token)

	arr, err2 := bot.GetUpdates(&GetUpdatesOpt{Limit: 1})
	assert.Nil(t, err2, "GetUpdates err")

	if len(arr) < 1 {
		return
	}

	up := arr[0]

	assert.NotNil(t, up.Message, "Message present")
	chatID := up.Message.From.ID

	assert.False(t, up.Message.From.IsBot, "IsBot")
	assert.True(t, chatID > 0, "Positive chatID")
}
