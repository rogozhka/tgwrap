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

//
// ParseModes represents formatting optiona for SendMessage()
//
// The Bot API supports basic formatting for messages.
// You can use bold and italic text, as well as inline links
// and pre-formatted code in your bots' messages.
// Telegram clients will render them accordingly.
// You can use either markdown-style or HTML-style formatting.
//
// Note that Telegram clients will display an alert to the user
// before opening an inline link (‘Open this link?’ together with the full URL).
//
// Links 'tg://user?id=<user_id>' can be used to mention
// a user by their id without using a username.
// Please note:  These links will work only if they are used
// inside an inline link. These mentions are only guaranteed
// to work if the user has contacted the bot in the past
// or is a member in the group where he was mentioned.
//
type ParseModes string

const (

	//
	// ParseModeDefault is simple text by default
	//
	ParseModeDefault ParseModes = ""

	//
	// ParseModeHTML The following tags are currently supported:
	//
	//  <b>bold</b>, <strong>bold</strong>
	//  <i>italic</i>, <em>italic</em>
	//  <a href="http://www.example.com/">inline URL</a>
	//  <a href="tg://user?id=123456789">inline mention of a user</a>
	//  <code>inline fixed-width code</code>
	//  <pre>pre-formatted fixed-width code block</pre>
	//
	ParseModeHTML ParseModes = "HTML"

	//
	// ParseModeMarkdown Use the following syntax in your message:
	//
	//  *bold text*
	//  _italic text_
	//  [inline URL](http://www.example.com/)
	//  [inline mention of a user](tg://user?id=123456789)
	//  `inline fixed-width code`
	//  ```block_language
	//  pre-formatted fixed-width code block
	//  ```
	//
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
