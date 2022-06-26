package tgwrap

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewBotWithClientAndURL validates variables passed.
func TestNewBotWithClientAndURL(t *testing.T) {
	is := assert.New(t)

	strToken := "dummy-token"
	client := &http.Client{}
	apiURL := "http://nonexistent-host.org/withPath"

	bot := NewBotWithClientAndURL(strToken, client, apiURL)

	is.NotNil(bot, "NewBotWithClientAndURL")
	is.Equal(strToken, bot.token, "token")
	is.Equal(client, bot.client, "client")
	is.Equal(apiURL, bot.apiURL, "apiURL")
}

// TestNewBot validates variables passed.
func TestNewBot(t *testing.T) {
	is := assert.New(t)

	strToken := "dummy-test-token"
	bot := NewBot(strToken)

	is.NotNil(bot, "NewBot")
	is.Equal(strToken, bot.token, "token")
	is.Equal(TelegramBotAPI, bot.apiURL, "apiURL")
}

// TestNewBotWithClient validates variables passed.
func TestNewBotWithClient(t *testing.T) {
	is := assert.New(t)

	strToken := "dummy-test-token"
	client := &http.Client{}

	bot := NewBotWithClient(strToken, client)

	is.NotNil(bot, "NewBotWithClient")
	is.Equal(strToken, bot.token, "token")
	is.Equal(client, bot.client, "client")
}

// TestNewBotWithClient validates variables passed.
func TestNewBotWithClient_nil(t *testing.T) {
	is := assert.New(t)

	strToken := "dummy-test-token"

	bot := NewBotWithClient(strToken, nil)

	is.NotNil(bot, "NewBotWithClient")
	is.Equal(strToken, bot.token, "token")
	is.NotNil(bot.client, "client should be created internally")
}
