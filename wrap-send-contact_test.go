package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendContactWrap(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)

	firstName := "CIA"
	phoneNumber := "(703) 482-0623"

	_, err := bot.SendContact(chatID, phoneNumber, firstName,
		&SendContactOpt{
			LastName: "Official",
		})

	is.Nil(err, "SendContact err")
}
