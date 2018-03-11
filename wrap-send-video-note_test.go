package tgwrap

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func Test_SendVideoNoteWrap(t *testing.T) {
	token := getTokenEnv()
	bot := createBot(token)

	envFilePath := "TGWRAP_TEST_VIDEO_PATH"
	filePath := os.Getenv(envFilePath)
	if len(filePath) < 1 {
		return
	}

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	inputFile := NewInputFileLocal(filePath)

	_, err4 := bot.SendVideoNote(chatID, inputFile,
		&SendVideoNoteOpt{
			DisableNotification: false,
		})

	assert.Nil(t, err4, "SendMessage err")
}
