package tgwrap

import (
	"context"
	"fmt"
	"reflect"

	"github.com/rogozhka/thestruct"
)

// SendVoiceOpt represents optional params for SendVoice.
type SendVoiceOpt struct {
	commonRequestOptions

	// Caption is audio caption, 0-200 characters.
	Caption string `json:"caption,omitempty"`

	// Duration of the audio in seconds.
	Duration uint `json:"duration,omitempty"`

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

// SendVoice used to send audio files, if you want Telegram clients to display the file
// as a playable voice message. For this to work, your audio must be in an .ogg file
// encoded with OPUS (other formats may be sent as Audio or Document).
// On success, the sent Message is returned. Bots can currently send voice messages
// of up to 50 MB in size, this limit may be changed in the future.
//
// chatID: (int64 or string) is unique identifier for the target chat
// or username of the target channel (in the format @channelusername).
//
// voice: (*InputFile or string) Audio to send. Pass a file_id as String to send
// an audio that exists on the Telegram servers (recommended), pass an HTTP URL as a String
// for Telegram to get an audio from the Internet, or upload a new file using multipart/form-data.
// using &NewInputFileLocal("<file path>").
//
// opt: (can be nil) optional params.
func (p *bot) SendVoice(chatID interface{}, voice interface{}, opt *SendVoiceOpt) (*Message, error) {

	type sendFormat struct {
		ChatID       string `json:"chat_id"`
		SendVoiceOpt `json:",omitempty"`
		// Voice is Audio file to send. Pass a file_id as String to send a photo that exists
		// on the Telegram servers (recommended), pass an HTTP URL as a String
		// for Telegram to get a photo from the Internet,
		// or upload a new photo using multipart/form-data.
		//
		// InputFile should have MarshalText interface
		Voice interface{} `json:"voice" form:"file"`
	}
	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		Voice:  voice,
	}

	if opt == nil {
		opt = &SendVoiceOpt{}
	}
	if opt.Context == nil {
		opt.Context = context.Background()
	}
	dataSend.SendVoiceOpt = *opt

	var resp struct {
		GenericResponse
		Result *Message `json:"result"`
	}
	sender := p.sendJSON
	tt := thestruct.Type(reflect.TypeOf(voice))
	if "InputFileLocal" == tt.Name() {
		sender = p.sendFormData
	}
	err := p.getAPIResponse(opt.Context, "sendVoice", sender, dataSend, &resp)
	return resp.Result, err
}
