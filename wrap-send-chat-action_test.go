package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendChatAction(t *testing.T) {
	token := getTokenEnv()
	bot := createBot(token)
	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	b2, err4 := bot.SendChatAction(chatID, ChatActionUploadPhoto)
	assert.Nil(t, err4, "SendChatAction err")
	assert.True(t, b2, "ChatActionTyping")
}
