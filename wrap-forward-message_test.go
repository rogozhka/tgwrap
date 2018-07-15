package tgwrap

import (
	"testing"
	"github.com/SergeyDonskoy/tgwrap/config"
	"github.com/stretchr/testify/assert"
)

func Test_ForwardMessageWrap(t *testing.T) {

	bot := NewBot(config.Token)

	_, err := bot.ForwardMessage(config.ChatID, config.FromChatID, false, config.MessageID)
	assert.Nil(t, err)
}

