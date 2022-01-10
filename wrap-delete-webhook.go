package tgwrap

// DeleteWebhookOpt represents optional params for DeleteWebhook.
type DeleteWebhookOpt struct {
	commonRequestOptions

	// DropPendingUpdates is optional. Pass True to drop all pending updates.
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// DeleteWebhook is used to remove webhook integration if you decide
// to switch back to getUpdates. Returns True on success.
func (p *bot) DeleteWebhook(opt *DeleteWebhookOpt) (bool, error) {
	type sendFormat struct {
		DeleteWebhookOpt // optional part
	}
	dataSend := sendFormat{}
	if opt != nil {
		dataSend.DeleteWebhookOpt = *opt
	}
	var resp struct {
		GenericResponse
		Result bool `json:"result"`
	}
	err := p.getAPIResponse(opt.Context, "deleteWebhook", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
