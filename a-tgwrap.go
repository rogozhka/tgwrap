package tgwrap

import (
	"fmt"
	"net/http"
	"strings"
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

type BotOpt struct {

	//
	// Client is optional http.Client
	// used default if nil
	//
	Client *http.Client

	//
	// API is optional apiURL
	// used default if ""
	//
	API string

	//
	// HideToken if set,
	// real apiKey should be excluded from logs
	//
	HideToken bool
}

func optRefined(opt *BotOpt) *BotOpt {

	if nil == opt {
		opt = &BotOpt{}
	}

	if len(opt.API) < 1 {
		opt.API = TelegramBotAPI
	}

	if nil == opt.Client {
		opt.Client = &http.Client{
			Timeout: DefaultClientTimeout,
		}
	}

	return opt
}

type bot struct {
	token string

	client *http.Client

	apiURL string

	hideToken bool
}

//
// NewBot creates new object with associated token inside
//
func NewBot(token string) *bot {
	return NewBotWithOpt(token, nil)
}

//
// NewBotWithClient creates bot w/ associated token and optional HTTP client
//
func NewBotWithClient(token string, client *http.Client) *bot {
	return NewBotWithOpt(
		token,
		&BotOpt{
			Client: client,
		},
	)
}

//
// NewBotWithClientAndURL creates bot w/ token, client and different API URL (proxy)
//
func NewBotWithClientAndURL(token string, client *http.Client, apiURL string) *bot {
	return NewBotWithOpt(
		token,
		&BotOpt{
			Client: client,
			API:    apiURL,
		},
	)
}

//
// NewBotWithOpt is general method
// for creating bot w/ options
//
func NewBotWithOpt(token string, optRaw *BotOpt) *bot {

	opt := optRefined(optRaw)

	p := &bot{
		token:     strings.TrimSpace(token),
		client:    opt.Client,
		apiURL:    opt.API,
		hideToken: opt.HideToken,
	}

	return p
}

func (p *bot) maskToken(err error) error {

	if !p.hideToken {
		return err
	}

	return fmt.Errorf(strings.Replace(err.Error(), p.token, "Token", -1))
}
