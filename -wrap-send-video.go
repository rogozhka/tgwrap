package tgwrap

import (
	"fmt"
	"reflect"

	"github.com/rogozhka/tgwrap/internal/thestruct"
)

type SendVideoOpt struct {

	//
	// Video caption (may also be used when resending photos by file_id),
	// 0-200 characters
	//
	Caption string `json:"caption,omitempty"`

	//
	// Duration of sent video in seconds
	//
	Duration uint `json:"duration,omitempty"`

	//
	// Video width
	//
	Width uint `json:"width,omitempty"`

	//
	// Video height
	//
	Height uint `json:"height,omitempty"`
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

func (p *bot) SendVideo(chatID interface{}, video interface{}, opt *SendVideoOpt) (*Message, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		SendVideoOpt `json:",omitempty"`

		//
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

	var sender fCommandSender = p.sendJSON

	tt := thestruct.Type(reflect.TypeOf(video))
	if "InputFile" == tt.Name() && len(video.(*InputFile).Name()) > 0 {
		sender = p.sendFormData
	}

	err := p.getAPIResponse("sendVideo", sender, dataSend, &resp)
	if err != nil {
		return nil, fmt.Errorf("getAPIResponse ERROR:%v", err)
	}

	return resp.Result, nil
}
