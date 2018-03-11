package tgwrap

import (
	"fmt"
)

// UnbanChatMemberis used to unban a previously kicked user in a supergroup or channel.
// The user will not return to the group or channel automatically, but will be able to join via link, etc.
// The bot must be an administrator for this to work.
// Returns True on success.
//
// chatID: (uint64 or string) Unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// userID: (int64) Unique identifier of the target user
//
func (p *bot) UnbanChatMember(chatID interface{}, userID uint64) (bool, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		UserID uint64 `json:"user_id"`
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		UserID: userID,
	}

	var resp struct {
		GenericResponse

		Result bool `json:"result"`
	}

	var sender fCommandSender = p.sendJSON

	err := p.getAPIResponse("unbanChatMember", sender, dataSend, &resp)
	if err != nil {
		return false, fmt.Errorf("getAPIResponse ERROR:%v", err)
	}

	return resp.Result, nil
}
