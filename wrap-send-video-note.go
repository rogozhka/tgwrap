package tgwrap

import (
	"context"
	"fmt"
	"reflect"

	"github.com/rogozhka/thestruct"
)

// SendVideoNoteOpt represents optional params for SendVideoNote.
type SendVideoNoteOpt struct {
	commonRequestOptions

	// Duration of sent video in seconds.
	Duration uint `json:"duration,omitempty"`

	// Length is optional video length in bytes.
	Length uint `json:"width,omitempty"`

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

// SendVideoNote is used to send video notes.
// As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long.
// Use this method to send video messages. On success, the sent Message is returned.
//
// chatID: (int64 or string) is unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// video: (*InputFile or string) Video note to send.
// Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or
// upload a new video using multipart/form-data.
// Sending video notes by a URL is currently unsupported
//
// opt: (can be nil) optional params
func (p *bot) SendVideoNote(chatID interface{}, video interface{}, opt *SendVideoNoteOpt) (*Message, error) {

	type sendFormat struct {
		ChatID           string `json:"chat_id"`
		SendVideoNoteOpt `json:",omitempty"`

		// VideoNote to send. Pass a file_id as String to send a photo that exists
		// on the Telegram servers (recommended), pass an HTTP URL as a String
		// for Telegram to get a photo from the Internet,
		// or upload a new photo using multipart/form-data.
		//
		// InputFile should have MarshalText interface
		//
		VideoNote interface{} `json:"video_note" form:"file"`
	}
	dataSend := sendFormat{
		ChatID:    fmt.Sprint(chatID),
		VideoNote: video,
	}

	if opt == nil {
		opt = &SendVideoNoteOpt{}
	}
	if opt.Context == nil {
		opt.Context = context.Background()
	}
	dataSend.SendVideoNoteOpt = *opt

	var resp struct {
		GenericResponse
		Result *Message `json:"result"`
	}
	sender := p.sendJSON
	tt := thestruct.Type(reflect.TypeOf(video))
	if "InputFileLocal" == tt.Name() {
		sender = p.sendFormData
	}
	err := p.getAPIResponse(opt.Context, "sendVideoNote", sender, dataSend, &resp)
	return resp.Result, err
}
