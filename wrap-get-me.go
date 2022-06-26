package tgwrap

import "context"

// GetMe is used for testing your bot's auth token.
// Requires no parameters. Returns basic information
// about the bot in form of a User object.Telegram API method.
func (p *bot) GetMe() (*User, error) {
	return p.GetMeContext(nil)
}

func (p *bot) GetMeContext(ctx context.Context) (*User, error) {
	var resp struct {
		GenericResponse
		Result *User `json:"result"`
	}
	err := p.getAPIResponse(ctx, "getMe", p.sendJSON, nil, &resp)
	return resp.Result, err
}
