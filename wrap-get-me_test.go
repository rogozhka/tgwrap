package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMe(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	u, err2 := bot.GetMe()
	is.Nil(err2, "GetMe err")
	is.NotNil(u, "GetMe User")

	is.True(u.IsBot, "IsBot")
	is.True(u.ID > 0, "Positive userID")
}
