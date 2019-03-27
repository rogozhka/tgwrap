package tgwrap

import (
	"net/http"
	"time"
)

//
// TelegramBotAPI is default api address
// used in default client
//
const TelegramBotAPI = "https://api.telegram.org/bot"

//
// DefaultClientTimeout is waiting for result interval
// used by default in default client
//
const DefaultClientTimeout = time.Second * 50

type bot struct {
	token string

	client *http.Client

	apiURL string
}

//
// NewBot creates new object with associated token inside
//
func NewBot(token string) *bot {
	return NewBotWithClient(token, &http.Client{
		Timeout: DefaultClientTimeout,
	})
}

//
// NewBotWithClient creates bot w/ associated token and optional HTTP client
//
func NewBotWithClient(token string, client *http.Client) *bot {
	p := &bot{
		token:  token,
		client: client,
		apiURL: TelegramBotAPI,
	}

	if nil == p.client {
		p.client = &http.Client{
			Timeout: DefaultClientTimeout,
		}
	}

	return p
}

//
// NewBotWithClientAndURL creates bot w/ token, client and different API URL (proxy)
//
func NewBotWithClientAndURL(token string, client *http.Client, apiURL string) *bot {

	p := NewBotWithClient(token, client)
	p.apiURL = apiURL

	return p
}
