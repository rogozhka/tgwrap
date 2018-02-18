package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendContactWrap(t *testing.T) {
	token := getTokenEnv()
	bot := createBot(token)
	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	firstName := "CIA"
	phoneNumber := "(703) 482-0623"

	_, err4 := bot.SendContact(chatID, phoneNumber, firstName,
		&SendContactOpt{
			LastName: "Official",
		})

	assert.Nil(t, err4, "SendContact err")
}
