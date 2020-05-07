package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendAudio(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")
	filePath := requireEnv("TGWRAP_TEST_AUDIO_PATH")

	inputFile := NewInputFileLocal(filePath)

	_, err := bot.SendAudio(chatID, inputFile,
		&SendAudioOpt{
			Caption:             "TGWRAP_TEST_AUDIO_PATH",
			Duration:            uint(3),
			DisableNotification: true,
		})

	assert.Nil(t, err, "SendAudio")
}
