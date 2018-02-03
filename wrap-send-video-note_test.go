package tgwrap

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func Test_SendVideoNoteWrap(t *testing.T) {
	token, err := getTokenEnv()
	assert.Nil(t, err, "Token")
	if len(token) < 1 {
		return
	}
	bot := NewBot(token)

	envFilePath := "TGWRAP_TEST_VIDEO_PATH"
	filePath := os.Getenv(envFilePath)
	if len(filePath) < 1 {
		return
		//panic(fmt.Errorf("Cannot find token env:%v", envFilePath))
	}

	arr, err2 := bot.GetUpdates(&GetUpdatesOpt{Limit: 1})
	assert.Nil(t, err2, "GetUpdates err")

	if len(arr) < 1 {
		return
	}

	up := arr[0]
	chatID := up.Message.From.ID

	inputFile := NewInputFileLocal(filePath)

	_, err4 := bot.SendVideoNote(chatID, inputFile,
		&SendVideoNoteOpt{
			DisableNotification: false,
		})

	assert.Nil(t, err4, "SendMessage err")
}
