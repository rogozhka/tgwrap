package tgwrap

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func Test_SendVideoNoteWrap(t *testing.T) {
	bot := createTestBotFromEnv()

	envFilePath := "TGWRAP_TEST_VIDEO_PATH"
	filePath := os.Getenv(envFilePath)
	if len(filePath) < 1 {
		return
	}

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	inputFile := NewInputFileLocal(filePath)

	r, err := bot.SendVideoNote(chatID, inputFile,
		&SendVideoNoteOpt{
			DisableNotification: false,
		})

	assert.Nil(t, err, "err")
	assert.NotNil(t, r, "result")
	assert.True(t, r.ID > 0, "positive message id")
}
