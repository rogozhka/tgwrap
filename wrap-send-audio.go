package tgwrap

import (
	"context"
	"fmt"
	"reflect"

	"github.com/rogozhka/thestruct"
)

// SendAudioOpt represents optional params for SendAudio.
type SendAudioOpt struct {
	commonRequestOptions

	// Caption is audio caption, 0-200 characters.
	Caption string `json:"caption,omitempty"`

	// Duration of the audio in seconds.
	Duration uint `json:"duration,omitempty"`

	// Performer of the track.
	Performer string `json:"performer,omitempty"`

	// Title is track name.
	Title string `json:"title,omitempty"`

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

// SendAudio is used to send audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .mp3 format. On success, the sent Message is returned. Bots can currently
// send audio files of up to 50 MB in size, this limit may be changed in the future.
//
// For sending voice messages, use the SendVoice method instead.
//
// chatID: (int64 or string) is unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// audio: (*InputFile or string) Audio to send. Pass a file_id as string to send
// an audio that exists on the Telegram servers (recommended), pass an HTTP URL as a string
// for Telegram to get an audio from the Internet, or upload a new file using multipart/form-data.
// using &NewInputFileLocal("<file path>")
//
// opt: (can be nil) optional params
func (p *bot) SendAudio(chatID interface{}, audio interface{}, opt *SendAudioOpt) (*Message, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		SendAudioOpt `json:",omitempty"`

		// Audio to send. Pass a file_id as String to send a photo that exists
		// on the Telegram servers (recommended), pass an HTTP URL as a String
		// for Telegram to get a photo from the Internet,
		// or upload a new photo using multipart/form-data.
		//
		// InputFile should have MarshalText interface
		//
		Audio interface{} `json:"audio" form:"file"`
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		Audio:  audio,
	}
	if opt == nil {
		opt = &SendAudioOpt{}
		opt.Context = context.Background()
	} else {
		dataSend.SendAudioOpt = *opt
	}
	var resp struct {
		GenericResponse
		Result *Message `json:"result"`
	}
	sender := p.sendJSON
	tt := thestruct.Type(reflect.TypeOf(audio))
	if "InputFileLocal" == tt.Name() {
		sender = p.sendFormData
	}
	err := p.getAPIResponse(opt.Context, "sendAudio", sender, dataSend, &resp)
	return resp.Result, err
}
