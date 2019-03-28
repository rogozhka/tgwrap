package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendMessage(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv("TGWRAP_TEST_PERSONAL_CHAT_ID")

	m, err := bot.SendMessage(chatID, "test message ;)",
		&SendMessageOpt{
			DisableNotification: true,
		})

	assert.Nil(t, err, "SendMessage err")
	assert.NotNil(t, m, "Message present")
}
