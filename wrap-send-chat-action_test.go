package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendChatAction(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)

	b2, err := bot.SendChatAction(chatID, ChatActionUploadPhoto)
	is.Nil(err, "SendChatAction err")
	is.True(b2, "ChatActionTyping")
}
