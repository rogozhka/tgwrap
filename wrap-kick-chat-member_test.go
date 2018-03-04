package tgwrap

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test_KickChatMember(t *testing.T) {
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
	fmt.Println(arr)
	up := arr[0]
	chatID := up.Message.From.ID
	userID := up.Message.From.ID
	opts := KickChatMemberOpt{UntilDate: 60}

	_, err4 := bot.KickChatMember(chatID, userID, &opts)
	assert.Nil(t, err4, "KickChatMember err")
}
