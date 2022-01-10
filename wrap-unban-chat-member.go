package tgwrap

import (
	"fmt"
)

// UnbanChatMember is used to unban a previously kicked user in a supergroup or channel.
// The user will not return to the group or channel automatically, but will be able to join via link, etc.
// The bot must be an administrator for this to work.
// Returns True on success.
//
// chatID: (int64 or string) is  unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
// userID is unique identifier of the target user.
func (p *bot) UnbanChatMember(chatID interface{}, userID int64) (bool, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`
		UserID int64  `json:"user_id"`
	}
	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		UserID: userID,
	}
	var resp struct {
		GenericResponse
		Result bool `json:"result"`
	}
	err := p.getAPIResponse(nil, "unbanChatMember", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
