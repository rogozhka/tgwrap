package tgwrap

import (
	"fmt"
)

//
// SendContactOpt represents optional params for SendContact
//
type SendContactOpt struct {
	commonRequestOptions

	//
	// Contact's last name
	//
	LastName string `json:"last_name,omitempty"`

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

//
// SendContact is used to to send phone contacts.
// On success, the sent Message is returned.
//
// chatID: (int64 or string) is unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// phoneNumber: (string) Contact's phone number.
//
// firstName: (string) Contact's first name.
//
// opt: (can be nil) optional params
//
func (p *bot) SendContact(chatID interface{}, phoneNumber string, firstName string, opt *SendContactOpt) (*Message, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		PhoneNumber string `json:"phone_number"`

		FirstName string `json:"first_name"`

		SendContactOpt `json:",omitempty"`
	}

	dataSend := sendFormat{
		ChatID:      fmt.Sprint(chatID),
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}

	if opt != nil {
		dataSend.SendContactOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result *Message `json:"result"`
	}

	err := p.getAPIResponse(opt.Context, "sendContact", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
