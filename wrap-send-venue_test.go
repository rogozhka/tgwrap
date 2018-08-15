package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendVenueWrap(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	longitude := 12.51
	latitude := 41.89
	title := "Rome"
	address := "Center"

	_, err4 := bot.SendVenue(chatID, latitude, longitude, title, address,
		&SendVenueOpt{
			DisableNotification: false,
		})

	assert.Nil(t, err4, "SendMessage err")
}
