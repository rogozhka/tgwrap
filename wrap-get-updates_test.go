package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUpdates(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	arr, err2 := bot.GetUpdates(&GetUpdatesOpt{Limit: 1})
	is.Nil(err2, "GetUpdates err")

	if len(arr) < 1 {
		return
	}

	up := arr[0]

	is.NotNil(up.Message, "Message present")
	chatID := up.Message.From.ID

	//	is.False(up.Message.From.IsBot, "IsBot")
	is.True(chatID > 0, "Positive chatID")
}
