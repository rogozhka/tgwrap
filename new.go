package tgwrap

import (
	"net/http"
	"time"
)

const telegramBotAPI = "https://api.telegram.org/bot"

type bot struct {
	token string

	client *http.Client

	apiURL string
}

//
// NewBot creates new object with associated token inside
//
// note: object is object, new is new
// and this is stupid comment to make linter happy
//
func NewBot(token string) IBot {
	return createBot(token)
}

func NewBotWithClient(token string, client *http.Client) IBot {
	return createBotWithClient(token, client)
}

func NewBotWithClientAndURL(token string, client *http.Client, apiURL string) IBot {
	return createBotWithClientAndURL(token, client, apiURL)
}

//
// createBot is used for testing purposes
// to test new methods still unexposed in IBot
//
func createBot(token string) *bot {
	return createBotWithClient(token, &http.Client{
		Timeout: time.Second * 10,
	})
}

func createBotWithClient(token string, client *http.Client) *bot {
	return &bot{
		token:  token,
		client: client,
		apiURL: telegramBotAPI,
	}
}

func createBotWithClientAndURL(token string, client *http.Client, apiURL string) *bot {
	return &bot{
		token:  token,
		client: client,
		apiURL: apiURL,
	}
}
