package tgwrap

import (
	"fmt"
)

// ChatActions type defines string arguments range for SendChatAction.
type ChatActions string

const (

	// ChatActionTyping indicates bot is typing (this looks frighteningly sometimes)
	ChatActionTyping ChatActions = "typing"

	// ChatActionUploadPhoto indicates bot is uploading photo
	ChatActionUploadPhoto ChatActions = "upload_photo"

	// ChatActionRecordVideo indicates bot is recording video
	ChatActionRecordVideo ChatActions = "record_video"

	// ChatActionRecordAudio indicates bot is recording audio
	ChatActionRecordAudio ChatActions = "record_audio"

	// ChatActionUploadDocument indicates bot is uploading document
	ChatActionUploadDocument ChatActions = "upload_document"

	// ChatActionFindLocation indicates bot is in search
	// for location (is it really working on any tg client?)
	ChatActionFindLocation ChatActions = "find_location"

	// ChatActionRecordVideoNote indicates bot is recording video note
	ChatActionRecordVideoNote ChatActions = "record_video_note"

	// ChatActionUploadVideoNote indicates bot is uploading note
	ChatActionUploadVideoNote ChatActions = "upload_video_note"
)

// SendChatAction is used when you need to tell the user
// that something is happening on the bot's side. The status
// is set for 5 seconds or less (when a message arrives from
// your bot, Telegram clients clear its typing status).
// Returns True on success.
//
// chatID: unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// action: type of action to broadcast.
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

	err := p.getAPIResponse(nil, "sendChatAction", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
