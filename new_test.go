package tgwrap

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// Test_createBotWithClientAndURL validates variables passed
//
func Test_createBotWithClientAndURL(t *testing.T) {

	strToken := "dummy-token"
	client := &http.Client{}
	apiURL := "http://nonexistent-host.org/withPath"

	bot := createBotWithClientAndURL(strToken, client, apiURL)

	assert.NotNil(t, bot, "createBotWithClientAndURL")
	assert.Equal(t, strToken, bot.token, "token")
	assert.Equal(t, client, bot.client, "client")
	assert.Equal(t, apiURL, bot.apiURL, "apiURL")
}
