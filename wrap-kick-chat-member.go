package tgwrap

import (
	"fmt"
)

//
// KickChatMemberOpt represents optional params for KickChatMember()
//
type KickChatMemberOpt struct {

	//
	// Date when the user will be unbanned, unix time.
	// If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
	//
	UntilDate int64
}

// KickChatMember is used to kick a user from a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// Returns True on success.
//
// chatID: (uint64 or string) Unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// userID: (int64) Unique identifier of the target user
//
// opt: (can be nil) optional params
//
func (p *bot) KickChatMember(chatID interface{}, userID uint64, opt *KickChatMemberOpt) (bool, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		UserID uint64 `json:"user_id"`

		KickChatMemberOpt `json:",omitempty"`
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		UserID: userID,
	}

	if opt != nil {
		dataSend.KickChatMemberOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result bool `json:"result"`
	}

	var sender fCommandSender = p.sendJSON

	err := p.getAPIResponse("kickChatMember", sender, dataSend, &resp)
	if err != nil {
		return false, fmt.Errorf("getAPIResponse ERROR:%v", err)
	}

	return resp.Result, nil
}
