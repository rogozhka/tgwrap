package tgwrap

import (
	"fmt"
	"reflect"

	"github.com/rogozhka/tgwrap/internal/thestruct"
)

type SendVoiceOpt struct {

	//
	// Audio caption, 0-200 characters
	// 0-200 characters
	//
	Caption string `json:"caption,omitempty"`

	//
	// Duration of the audio in seconds
	//
	Duration uint `json:"duration,omitempty"`

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
// SendVoice to send audio files, if you want Telegram clients to display the file
// as a playable voice message. For this to work, your audio must be in an .ogg file
// encoded with OPUS (other formats may be sent as Audio or Document).
// On success, the sent Message is returned. Bots can currently send voice messages
// of up to 50 MB in size, this limit may be changed in the future.
//
func (p *bot) SendVoice(chatID interface{}, voice interface{}, opt *SendVoiceOpt) (*Message, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		SendVoiceOpt `json:",omitempty"`

		//
		// Audio file to send. Pass a file_id as String to send a photo that exists
		// on the Telegram servers (recommended), pass an HTTP URL as a String
		// for Telegram to get a photo from the Internet,
		// or upload a new photo using multipart/form-data.
		//
		// InputFile should have MarshalText interface
		//
		Voice interface{} `json:"voice" form:"file"`
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		Voice:  voice,
	}

	if opt != nil {
		dataSend.SendVoiceOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result *Message `json:"result"`
	}

	var sender fCommandSender = p.sendJSON

	tt := thestruct.Type(reflect.TypeOf(voice))
	if "InputFile" == tt.Name() && len(voice.(*InputFile).Name()) > 0 {
		sender = p.sendFormData
	}

	err := p.getAPIResponse("sendVoice", sender, dataSend, &resp)
	if err != nil {
		return nil, fmt.Errorf("getAPIResponse ERROR:%v", err)
	}

	return resp.Result, nil
}
