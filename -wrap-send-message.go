package tgwrap

import (
	"fmt"
)

//
// SendMessageOpt represents optional params for SendMessage
//
type SendMessageOpt struct {
	//
	// Send Markdown or HTML, if you want Telegram apps
	// to show bold, italic, fixed-width text or inline URLs
	// in your bot's message.
	//
	ParseMode ParseModes `json:"parse_mode,omitempty"`

	//
	// Disables link previews for links in this message
	//
	DisableWebPreview bool `json:"disable_web_page_preview,omitempty"`

	//
	// Sends the message silently. Users will receive a notification with no sound.
	//
	DisableNotification bool `json:"disable_notification,omitempty"`

	//
	// If the message is a reply, ID of the original message
	//
	ReplyToID uint64 `json:"reply_to_message_id,omitempty"`

	//
	// Additional interface options. A JSON-serialized object
	// for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard
	// or to force a reply from the user.
	//
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type ParseModes string

const (
	ParseModeDefault  ParseModes = ""
	ParseModeHTML     ParseModes = "HTML"
	ParseModeMarkdown ParseModes = "Markdown"
)

//
// SendMessage is used method to send text messages.
//
// chatID: unique identifier for the target chat
// or username(!) of the target channel (in the format @channelusername)
//
// text: text of the message to be sent
//
// opt: (can be nil) optional params
//
func (p *bot) SendMessage(chatID interface{}, text string, opt *SendMessageOpt) (*Message, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`
		Text   string `json:"text"`

		SendMessageOpt // optional part
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID), // don't care about checking fmt, Telegram will response with error if invalid ID
		Text:   text,
	}

	if opt != nil {
		dataSend.SendMessageOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result *Message `json:"result"`
	}

	err := p.getAPIResponse("sendMessage", p.sendJSON, dataSend, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
