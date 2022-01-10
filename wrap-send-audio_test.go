package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendAudio(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)
	filePath := requireEnv(envTestAudioPath)

	inputFile := NewInputFileLocal(filePath)

	_, err := bot.SendAudio(chatID, inputFile,
		&SendAudioOpt{
			Caption:             envTestAudioPath,
			Duration:            uint(3),
			DisableNotification: true,
		})

	assert.Nil(t, err, "SendAudio")
}
