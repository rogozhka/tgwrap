package tgwrap

import "fmt"

type SetWebhookOpt struct {
	Certificate    *InputFile `json:"certificate,omitempty"`
	MaxConnections uint       `json:"max_connections"`
	AllowedUpdates []string   `json:"allowed_updates"`
}

func (p *bot) SetWebhook(url string, opt *SetWebhookOpt) (bool, error) {

	return false, fmt.Errorf("Not implemented yet")
}
