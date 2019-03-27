package tgwrap

//
// StopMessageLiveLocationOpt represents optional params for StopMessageLiveLocation
//
type StopMessageLiveLocationOpt struct {

	//
	// Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	//
	ChatID interface{} `json:"chat_id,omitempty"`

	//
	// Required if inline_message_id is not specified.
	// Identifier of the sent message
	//
	MessageID uint `json:"message_id,omitempty"`

	//
	// Required if chat_id and message_id are not specified.
	// Identifier of the inline message
	//
	InlineMessageID string `json:"inline_message_id,omitempty"`

	//
	// Additional interface options. A JSON-serialized object
	// for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard
	// or to force a reply from the user.
	//
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

//
// StopMessageLiveLocation is used to stop updating a live location message sent by the bot or via the bot (for inline bots) before live_period expires.
// On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
//
// opt: (can be nil) optional params
//
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

	err := p.getAPIResponse("stopMessageLiveLocation", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
