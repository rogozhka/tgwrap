package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendMessage(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestPersonalChatID)

	m, err := bot.SendMessage(chatID, "test message ;)",
		&SendMessageOpt{
			DisableNotification: true,
		})

	is.Nil(err, "SendMessage err")
	is.NotNil(m, "Message present")
}
