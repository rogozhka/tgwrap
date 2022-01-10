package tgwrap

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func Test_SendVideoNoteWrap(t *testing.T) {
	bot := createTestBotFromEnv()

	filePath := os.Getenv(envTestVideoPath)
	if len(filePath) < 1 {
		return
	}

	chatID := requireEnv(envTestChatID)

	inputFile := NewInputFileLocal(filePath)

	r, err := bot.SendVideoNote(chatID, inputFile,
		&SendVideoNoteOpt{
			DisableNotification: false,
		})

	assert.Nil(t, err, "SendVideoNote err")
	assert.NotNil(t, r, "result")
	assert.True(t, r.ID > 0, "positive message id")
}
