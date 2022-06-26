package tgwrap

import "context"

// GetWebhookInfo is used to get current webhook status.
// If the bot is using getUpdates, will return an object
// with the url field empty.
func (p *bot) GetWebhookInfo() (*WebhookInfo, error) {
	return p.GetWebhookInfoContext(nil)
}

func (p *bot) GetWebhookInfoContext(ctx context.Context) (*WebhookInfo, error) {
	var resp struct {
		GenericResponse
		Result *WebhookInfo `json:"result"`
	}
	err := p.getAPIResponse(ctx, "getWebhookInfo", p.sendJSON, nil, &resp)
	return resp.Result, err
}
