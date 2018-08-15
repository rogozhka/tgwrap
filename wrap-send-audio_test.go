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

	_, err4 := bot.SendAudio(chatID, inputFile,
		&SendAudioOpt{
			DisableNotification: true,
			Duration:            uint(3),
			Caption:             "TGWRAP_TEST_AUDIO_PATH",
		})

	assert.Nil(t, err4, "SendAudio")
}
