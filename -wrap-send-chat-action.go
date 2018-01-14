package tgwrap

import (
	"fmt"
)

type ChatActions string

const (
	ChatActionTyping          ChatActions = "typing"
	ChatActionUploadPhoto     ChatActions = "upload_photo"
	ChatActionRecordVideo     ChatActions = "record_video"
	ChatActionRecordAudio     ChatActions = "record_audio"
	ChatActionUploadDocument  ChatActions = "upload_document"
	ChatActionFindLocation    ChatActions = "find_location"
	ChatActionRecordVideoNote ChatActions = "record_video_note"
	ChatActionUploadVideoNote ChatActions = "upload_video_note"
)

func (p *bot) SendChatAction(chatID interface{}, action ChatActions) (bool, error) {

	type sendFormat struct {
		ChatID string      `json:"chat_id"`
		Action ChatActions `json:"action"`
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		Action: action,
	}

	var resp struct {
		GenericResponse

		Result bool `json:"result"`
	}

	err := p.getResponse("sendChatAction", p.sendJSON, dataSend, &resp)
	if err != nil {
		return false, err
	}

	return resp.Result, nil
}
