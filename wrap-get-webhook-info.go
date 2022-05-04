package tgwrap

// GetWebhookInfo is used to get current webhook status.
// If the bot is using getUpdates, will return an object
// with the url field empty.
func (p *bot) GetWebhookInfo() (*WebhookInfo, error) {
	var resp struct {
		GenericResponse
		Result *WebhookInfo `json:"result"`
	}
	err := p.getAPIResponse(nil, "getWebhookInfo", p.sendJSON, nil, &resp)
	return resp.Result, err
}
