package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendChatAction(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	b2, err := bot.SendChatAction(chatID, ChatActionUploadPhoto)
	assert.Nil(t, err, "SendChatAction err")
	assert.True(t, b2, "ChatActionTyping")
}
