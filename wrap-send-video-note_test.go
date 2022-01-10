package tgwrap

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func Test_SendVideoNoteWrap(t *testing.T) {
	is := assert.New(t)
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
	is.Nil(err, "SendVideoNote err")
	is.NotNil(r, "result")
	is.True(r.ID > 0, "positive message id")
}
