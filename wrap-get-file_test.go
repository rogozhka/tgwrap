package tgwrap

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test_GetFileWrap(t *testing.T) {
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

	file, err4 := bot.GetFile("BAADAgADgAIAArpQwEga8pLP-1m7UgI")
	fmt.Print(file)
	assert.Nil(t, err4, "GetFile err")

}
