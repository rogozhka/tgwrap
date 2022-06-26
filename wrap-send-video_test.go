package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendVideoWrap(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)
	filePath := requireEnv(envTestVideoPath)

	inputFile := NewInputFileLocal(filePath)
	_, err := bot.SendVideo(chatID, inputFile,
		&SendVideoOpt{
			DisableNotification: true,
		})
	is.Nil(err, "SendVideo err")
}
