package tgwrap

// EditMessageLiveLocationOpt represents optional params for EditMessageLiveLocation.
type EditMessageLiveLocationOpt struct {
	commonRequestOptions

	// ChatID required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID interface{} `json:"chat_id,omitempty"`

	// MessageID required if inline_message_id is not specified.
	// Identifier of the sent message.
	MessageID int64 `json:"message_id,omitempty"`

	// InlineMessageID required if chat_id and message_id are not specified.
	// Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// ReplyMarkup - additional interface options. A JSON-serialized object
	// for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard
	// or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// EditMessageLiveLocation is used to edit live location messages sent by the bot or via the bot (for inline bots).
// A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation.
// On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
//
// latitude: (float64) Latitude of new location.
//
// longitude: (float64) Longitude of new location.
//
// opt: (can be nil) optional params
func (p *bot) EditMessageLiveLocation(latitude float64, longitude float64, opt *EditMessageLiveLocationOpt) (interface{}, error) {

	type sendFormat struct {
		EditMessageLiveLocationOpt `json:",omitempty"`

		Latitude float64 `json:"latitude"`

		Longitude float64 `json:"longitude"`
	}

	dataSend := sendFormat{
		Latitude:  latitude,
		Longitude: longitude,
	}

	if opt != nil {
		dataSend.EditMessageLiveLocationOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result *Message `json:"result"`
	}

	err := p.getAPIResponse(opt.Context, "editMessageLiveLocation", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
