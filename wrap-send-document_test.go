package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendDocument(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	inputFile := NewInputFileLocal("wrap-send-document.go")

	_, err4 := bot.SendDocument(chatID, inputFile,
		&SendDocumentOpt{
			DisableNotification: true,
		})

	assert.Nil(t, err4, "SendMessage err")
}
