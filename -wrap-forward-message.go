package tgwrap

import (
	"fmt"
)

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

	var respStruct struct {
		GenericResponse

		Result *Message `json:"result"`
	}

	err := p.getResponse("forwardMessage", dataSend, &respStruct)
	if err != nil {
		return nil, err
	}

	return respStruct.Result, nil
}
