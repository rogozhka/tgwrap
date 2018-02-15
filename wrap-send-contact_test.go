package tgwrap

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_SendContactWrap(t *testing.T) {
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
	fmt.Println(up.Message.From.ID)
	firstName := "CIA"
	phoneNumber := "(703) 482-0623"

	_, err4 := bot.SendContact(chatID, phoneNumber, firstName,
		&SendContactOpt{
			LastName: "Official",
		})

	assert.Nil(t, err4, "SendContact err")
}
