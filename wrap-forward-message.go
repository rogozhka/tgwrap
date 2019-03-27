package tgwrap

import (
	"fmt"
)

//
// ForwardMessage is used to forward messages of any kind.
// On success, the sent Message is returned.
//
// chatID: unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// fromChatID: unique identifier for the chat
// where the original message was sent (or channel username
// in the format @channelusername)
//
// disableNotification: sends the message silently.
// Users will receive a notification with no sound.
//
// messageID: message identifier in the chat
// specified in fromChatID
//
func (p *bot) ForwardMessage(chatID interface{}, fromChatID interface{},
	disableNotification bool, messageID uint64) (*Message, error) {

	type sendFormat struct {
		ChatID              string `json:"chat_id"`
		FromChatID          string `json:"from_chat_id"`
		DisableNotification bool   `json:"disable_notification,omitempty"`
		MessageID           uint64 `json:"message_id"`
	}
	dataSend := sendFormat{
		MessageID:  messageID,
		ChatID:     fmt.Sprint(chatID),
		FromChatID: fmt.Sprint(fromChatID),
	}

	if disableNotification {
		dataSend.DisableNotification = true
	}

	var resp struct {
		GenericResponse

		Result *Message `json:"result"`
	}

	err := p.getAPIResponse("forwardMessage", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
