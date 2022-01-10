package tgwrap

import (
	"fmt"
	"reflect"

	"github.com/rogozhka/thestruct"
)

// SendVideoOpt represents optional params for SendVideo.
type SendVideoOpt struct {
	commonRequestOptions

	// Caption is video caption (may also be used when resending photos by file_id),
	// 0-200 characters
	Caption string `json:"caption,omitempty"`

	// Duration of sent video in seconds.
	Duration int `json:"duration,omitempty"`

	// Width is video width
	Width int `json:"width,omitempty"`

	// Height is video height.
	Height int `json:"height,omitempty"`

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

// SendVideo is used to send video files, Telegram clients support mp4 videos
// (other formats may be sent as Document). On success, the sent Message is returned.
// Bots can currently send video files of up to 50 MB in size,
// this limit may be changed in the future.
//
// chatID: (int64 or string) is unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// video: (*InputFile or string) Video to send. Pass a file_id as string to send
// a video that exists on the Telegram servers (recommended), pass an HTTP URL as a string
// for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data
// using &NewInputFileLocal("<file path>")
//
// opt: (can be nil) optional params
//
func (p *bot) SendVideo(chatID interface{}, video interface{}, opt *SendVideoOpt) (*Message, error) {

	type sendFormat struct {
		ChatID       string `json:"chat_id"`
		SendVideoOpt `json:",omitempty"`

		// Video to send. Pass a file_id as String to send a photo that exists
		// on the Telegram servers (recommended), pass an HTTP URL as a String
		// for Telegram to get a photo from the Internet,
		// or upload a new photo using multipart/form-data.
		//
		// InputFile should have MarshalText interface
		//
		Video interface{} `json:"video" form:"file"`
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		Video:  video,
	}
	if opt != nil {
		dataSend.SendVideoOpt = *opt
	}
	var resp struct {
		GenericResponse
		Result *Message `json:"result"`
	}
	sender := p.sendJSON
	tt := thestruct.Type(reflect.TypeOf(video))
	if "InputFileLocal" == tt.Name() {
		sender = p.sendFormData
	}
	err := p.getAPIResponse(opt.Context, "sendVideo", sender, dataSend, &resp)
	return resp.Result, err
}
