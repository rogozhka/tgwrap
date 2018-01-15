package tgwrap

import (
	"testing"

	"crypto/md5"
	"fmt"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_SendPhotoWrap(t *testing.T) {
	token, err := getTokenEnv()
	assert.Nil(t, err, "Token")
	if len(token) < 1 {
		return
	}

	bot := NewBot(token)

	arr, err2 := bot.GetUpdates(&GetUpdatesOpt{Limit: 1})
	assert.Nil(t, err2, "GetUpdates err")

	if len(arr) < 1 {
		return
	}

	up := arr[0]
	chatID := up.Message.From.ID

	photo := NewInputFileLocal("test/gray.png")

	_, err4 := bot.SendPhoto(chatID, photo,
		&SendPhotoOpt{
			DisableNotification: true,
			Caption:             fmt.Sprintf("Background: gray %v", md5.Sum([]byte(time.Now().String()))),
		})

	assert.Nil(t, err4, "SendMessage err")
}

func Test_SendPhotoURL(t *testing.T) {
	token, err := getTokenEnv()
	assert.Nil(t, err, "Token")
	if len(token) < 1 {
		return
	}

	bot := NewBot(token)

	arr, err2 := bot.GetUpdates(&GetUpdatesOpt{Limit: 1})
	assert.Nil(t, err2, "GetUpdates err")

	if len(arr) < 1 {
		return
	}

	up := arr[0]
	chatID := up.Message.From.ID

	photo := NewInputFileFromURL("https://raw.githubusercontent.com/rogozhka/tgwrap/master/test/rogozhka-digital.png")

	_, err4 := bot.SendPhoto(chatID, photo,
		&SendPhotoOpt{
			DisableNotification: true,
			Caption:             fmt.Sprintf("Background: rogozhka %v", md5.Sum([]byte(time.Now().String()))),
		})

	assert.Nil(t, err4, "SendMessage err")
}
