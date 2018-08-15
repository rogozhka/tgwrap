package tgwrap

import (
	"testing"

	"crypto/md5"
	"fmt"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_SendPhotoWrap(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	photo := NewInputFileLocal("test/gray.png")

	_, err4 := bot.SendPhoto(chatID, photo,
		&SendPhotoOpt{
			DisableNotification: true,
			Caption:             fmt.Sprintf("Background: gray %v", md5.Sum([]byte(time.Now().String()))),
		})

	assert.Nil(t, err4, "SendMessage err")
}

func Test_SendPhotoURL(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	photo := NewInputFileFromURL("https://raw.githubusercontent.com/rogozhka/tgwrap/master/test/rogozhka-digital.png")

	_, err4 := bot.SendPhoto(chatID, photo,
		&SendPhotoOpt{
			DisableNotification: true,
			Caption:             fmt.Sprintf("Background: rogozhka %v", md5.Sum([]byte(time.Now().String()))),
		})

	assert.Nil(t, err4, "SendMessage err")
}
