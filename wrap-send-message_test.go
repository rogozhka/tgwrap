package tgwrap

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test_SendMessage(t *testing.T) {
	token := getTokenEnv()
	bot := createBot(token)

	arr, err2 := bot.GetUpdates(&GetUpdatesOpt{Limit: 1})
	assert.Nil(t, err2, "GetUpdates err")

	if len(arr) < 1 {
		return
	}

	up := arr[0]
	chatID := up.Message.From.ID

	m, err4 := bot.SendMessage(chatID, fmt.Sprintf(";) %v", up.Message.Text),
		&SendMessageOpt{
			DisableNotification: true,
		})

	assert.Nil(t, err4, "SendMessage err")
	assert.NotNil(t, m, "Message present")
}
