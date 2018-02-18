package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendVideoWrap(t *testing.T) {
	token := getTokenEnv()
	bot := createBot(token)
	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")
	filePath := requireEnv("TGWRAP_TEST_VIDEO_PATH")

	inputFile := NewInputFileLocal(filePath)

	_, err4 := bot.SendVideo(chatID, inputFile,
		&SendVideoOpt{
			DisableNotification: true,
		})

	assert.Nil(t, err4, "SendMessage err")
}
