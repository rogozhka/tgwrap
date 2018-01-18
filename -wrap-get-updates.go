package tgwrap

//
// GetUpdatesOpt represents optional params for GetUpdates()
//
type GetUpdatesOpt struct {
	//
	// Identifier of the first update to be returned.
	// Must be greater by one than the highest among the identifiers
	// of previously received updates. By default, updates starting
	// with the earliest unconfirmed update are returned. An update
	// is considered confirmed as soon as getUpdates is called with
	// an offset higher than its update_id. The negative offset
	// can be specified to retrieve updates starting from -offset
	// update from the end of the updates queue.
	// All previous updates will forgotten.
	//
	Offset int64 `json:"offset,omitempty"`

	//
	// Limits the number of updates to be retrieved.
	// Values between 1—100 are accepted. Defaults to 100.
	//
	Limit uint `json:"limit,omitempty"`

	//
	// Timeout in seconds for long polling.
	// Defaults to 0, i.e. usual short polling. Should be positive,
	// short polling should be used for testing purposes only.
	//
	Timeout uint `json:"timeout,omitempty"`

	//
	// List the types of updates you want your bot to receive.
	// For example, specify [“message”, “edited_channel_post”, “callback_query”]
	// to only receive updates of these types. See Update for a complete list
	// of available update types. Specify an empty list to receive all updates
	// regardless of type (default). If not specified, the previous setting will be used.
	//
	// Please note that this parameter doesn't affect updates
	// created before the call to the getUpdates, so unwanted
	// updates may be received for a short period of time.
	//
	Allowed []string `json:"allowed_updates,omitempty"`
}

//
// GetUpdates is used to receive incoming updates using [long] polling
//
// opt: (can be nil) optional params
//
func (p *bot) GetUpdates(opt *GetUpdatesOpt) ([]*Update, error) {

	var arrRes []*Update

	type sendFormat struct {
		GetUpdatesOpt
	}

	dataSend := sendFormat{}

	if opt != nil {
		dataSend.GetUpdatesOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result []*Update `json:"result"`
	}

	err := p.getAPIResponse("getUpdates", p.sendJSON, dataSend, &resp)
	if err != nil {
		return arrRes, err
	}

	return resp.Result, nil
}
