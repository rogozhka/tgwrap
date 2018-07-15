package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/SergeyDonskoy/tgwrap/config"
)

func Test_getMe(t *testing.T) {

	bot := NewBot(config.Token)

	u, err2 := bot.GetMe()
	assert.Nil(t, err2, "GetMe err")
	assert.NotNil(t, u, "GetMe User")

	assert.True(t, u.IsBot, "IsBot")
	assert.True(t, u.ID > 0, "Positive userID")
}
