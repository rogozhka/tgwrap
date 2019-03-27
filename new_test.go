package tgwrap

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// Test_createBotWithClientAndURL validates variables passed
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
