package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendVenueWrap(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()
	chatID := requireEnv(envTestChatID)

	longitude := 12.51
	latitude := 41.89
	title := "Rome"
	address := "Center"
	_, err := bot.SendVenue(chatID, latitude, longitude, title, address,
		&SendVenueOpt{
			DisableNotification: false,
		})
	is.Nil(err, "SendVenue err")
}
