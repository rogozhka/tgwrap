package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMe(t *testing.T) {

	token, err := getTokenEnv()
	assert.Nil(t, err, "Token")
	if len(token) < 1 {
		return
	}

	bot := NewBot(token)

	u, err2 := bot.GetMe()
	assert.Nil(t, err2, "GetMe err")
	assert.NotNil(t, u, "GetMe User")

	assert.True(t, u.IsBot, "IsBot")
	assert.True(t, u.ID > 0, "Positive userID")
}
