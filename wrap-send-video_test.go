package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendVideoWrap(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")
	filePath := requireEnv("TGWRAP_TEST_VIDEO_PATH")

	inputFile := NewInputFileLocal(filePath)

	_, err := bot.SendVideo(chatID, inputFile,
		&SendVideoOpt{
			DisableNotification: true,
		})

	assert.Nil(t, err, "SendVideo err")
}
