package tgwrap

import (
	"context"
	"fmt"
	"reflect"

	"github.com/rogozhka/thestruct"
)

// SendDocumentOpt represents optional params for SendDocument.
type SendDocumentOpt struct {
	commonRequestOptions

	// Caption is document caption, 0-200 characters
	Caption string `json:"caption,omitempty"`

	// DisableNotification sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToID is optional ID of the original message if the message is a reply.
	ReplyToID int64 `json:"reply_to_message_id,omitempty"`

	// ReplyMarkup - additional interface options. A JSON-serialized object
	// for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard
	// or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// SendDocument to send general files. On success, the sent Message is returned.
// Bots can currently send files of any type of up to 50 MB in size, t
// his limit may be changed in the future.
//
// chatID: (int64 or string) is unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// document: (*InputFile or string) File to send. Pass a file_id as String to send
// a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String
// for Telegram to get an audio from the Internet, or upload a new file using multipart/form-data.
// using &NewInputFileLocal("<file path>")
//
// opt: (can be nil) optional params
//
func (p *bot) SendDocument(chatID interface{}, document interface{}, opt *SendDocumentOpt) (*Message, error) {

	type sendFormat struct {
		ChatID          string `json:"chat_id"`
		SendDocumentOpt `json:",omitempty"`
		Document        interface{} `json:"document" form:"file"`
	}
	dataSend := sendFormat{
		ChatID:   fmt.Sprint(chatID),
		Document: document,
	}

	if opt == nil {
		opt = &SendDocumentOpt{}
	}
	if opt.Context == nil {
		opt.Context = context.Background()
	}
	dataSend.SendDocumentOpt = *opt

	var resp struct {
		GenericResponse
		Result *Message `json:"result"`
	}
	sender := p.sendJSON
	tt := thestruct.Type(reflect.TypeOf(document))
	if "InputFileLocal" == tt.Name() {
		sender = p.sendFormData
	}
	err := p.getAPIResponse(opt.Context, "sendDocument", sender, dataSend, &resp)
	return resp.Result, err
}
