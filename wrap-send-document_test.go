package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendDocument(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)

	inputFile := NewInputFileLocal("wrap-send-document.go")

	_, err := bot.SendDocument(chatID, inputFile,
		&SendDocumentOpt{
			DisableNotification: true,
		})

	is.Nil(err, "SendDocument err")
}
