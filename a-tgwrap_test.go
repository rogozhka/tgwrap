package tgwrap

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// TestNewBotWithClientAndURL validates variables passed
//
func TestNewBotWithClientAndURL(t *testing.T) {

	strToken := "dummy-token"
	client := &http.Client{}
	apiURL := "http://nonexistent-host.org/withPath"

	bot := NewBotWithClientAndURL(strToken, client, apiURL)

	assert.NotNil(t, bot, "NewBotWithClientAndURL")
	assert.Equal(t, strToken, bot.token, "token")
	assert.Equal(t, client, bot.client, "client")
	assert.Equal(t, apiURL, bot.apiURL, "apiURL")
}

//
// TestNewBot validates variables passed
//
func TestNewBot(t *testing.T) {

	strToken := "dummy-test-token"
	bot := NewBot(strToken)

	assert.NotNil(t, bot, "NewBot")
	assert.Equal(t, strToken, bot.token, "token")
	assert.Equal(t, TelegramBotAPI, bot.apiURL, "apiURL")
}

//
// TestNewBotWithClient validates variables passed
//
func TestNewBotWithClient(t *testing.T) {

	strToken := "dummy-test-token"
	client := &http.Client{}

	bot := NewBotWithClient(strToken, client)

	assert.NotNil(t, bot, "NewBotWithClient")
	assert.Equal(t, strToken, bot.token, "token")
	assert.Equal(t, client, bot.client, "client")
}

//
// TestNewBotWithClient validates variables passed
//
func TestNewBotWithClient_nil(t *testing.T) {

	strToken := "dummy-test-token"

	bot := NewBotWithClient(strToken, nil)

	assert.NotNil(t, bot, "NewBotWithClient")
	assert.Equal(t, strToken, bot.token, "token")
	assert.NotNil(t, bot.client, "client should be created internally")
}
