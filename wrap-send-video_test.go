package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendVideoWrap(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)
	filePath := requireEnv(envTestVideoPath)

	inputFile := NewInputFileLocal(filePath)

	_, err := bot.SendVideo(chatID, inputFile,
		&SendVideoOpt{
			DisableNotification: true,
		})

	assert.Nil(t, err, "SendVideo err")
}
