package tgwrap

import (
	"fmt"
)

//
// SetChatTitle is used to change the title of a chat.
// Titles can't be changed for private chats.
// The bot must be an administrator in the chat for this to work
// and must have the appropriate admin rights. Returns True on success.
//
// Note: In regular groups (non-supergroups),
// this method will only work if the ‘All Members Are Admins’
// setting is off in the target group.
//
func (p *bot) SetChatTitle(chatID interface{}, title string) (bool, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`
		Title  string `json:"title"`
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		Title:  title,
	}

	var resp struct {
		GenericResponse

		Result bool `json:"result"`
	}

	err := p.getAPIResponse(nil, "setChatTitle", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
