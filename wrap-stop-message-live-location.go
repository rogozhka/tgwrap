package tgwrap

// StopMessageLiveLocationOpt for StopMessageLiveLocation optional params.
type StopMessageLiveLocationOpt struct {
	commonRequestOptions

	// ChatID is required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID interface{} `json:"chat_id,omitempty"`

	// MessageID is required if inline_message_id is not specified.
	// Identifier of the sent message
	MessageID int64 `json:"message_id,omitempty"`

	// InlineMessageID is required if ChatID and MessageID are not specified.
	// Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// ReplyMarkup - additional interface options. A JSON-serialized object
	// for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard
	// or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// StopMessageLiveLocation is used to stop updating a live location message sent by the bot or via the bot
// (for inline bots) before live_period expires.
// On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
func (p *bot) StopMessageLiveLocation(opt *StopMessageLiveLocationOpt) (interface{}, error) {

	type sendFormat struct {
		StopMessageLiveLocationOpt `json:",omitempty"`
	}
	dataSend := sendFormat{}
	if opt != nil {
		dataSend.StopMessageLiveLocationOpt = *opt
	}
	var resp struct {
		GenericResponse
		Result *Message `json:"result"`
	}
	err := p.getAPIResponse(opt.Context, "stopMessageLiveLocation", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
