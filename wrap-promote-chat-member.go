package tgwrap

import (
	"fmt"
)

//
// PromoteChatMemberOpt represents optional params for PromoteChatMemberOpt()
//
type PromoteChatMemberOpt struct {

	//
	// Pass True, if the administrator can change chat title, photo and other settings
	//
	CanChangeInfo bool `json:"can_change_info"`

	//
	// Pass True, if the administrator can create channel posts, channels only
	//
	CanPostMessages bool `json:"can_post_messages"`

	//
	// Pass True, if the administrator can edit messages of other users and can pin messages, channels only
	//
	CanEditMessages bool `json:"can_edit_messages"`

	//
	// Pass True, if the administrator can delete messages of other users
	//
	CanDeleteMessages bool `json:"can_delete_messages"`

	//
	// Pass True, if the administrator can invite new users to the chat
	//
	CanInviteUsers bool `json:"can_invite_users"`

	//
	// Pass True, if the administrator can restrict, ban or unban chat members
	//
	CanRestrictMembers bool `json:"can_restrict_members"`

	//
	// Pass True, if the administrator can pin messages, supergroups only
	//
	CanPinMessages bool `json:"can_pin_messages"`

	//
	// Pass True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted,
	// directly or indirectly (promoted by administrators that were appointed by him)
	//
	CanPromoteMembers bool `json:"can_promote_members"`
}

// PromoteChatMember is used to promote or demote a user in a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// Pass False for all boolean parameters to demote a user.
// Returns True on success.
//
// chatID: (uint64 or string) Unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// userID: (int64) Unique identifier of the target user
//
// opt: (can be nil) optional params
//
func (p *bot) PromoteChatMember(chatID interface{}, userID uint64, opt *PromoteChatMemberOpt) (bool, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		UserID uint64 `json:"user_id"`

		PromoteChatMemberOpt `json:",omitempty"`
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		UserID: userID,
	}

	if opt != nil {
		dataSend.PromoteChatMemberOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result bool `json:"result"`
	}

	var sender fCommandSender = p.sendJSON

	err := p.getAPIResponse("promoteChatMember", sender, dataSend, &resp)
	if err != nil {
		return false, fmt.Errorf("getAPIResponse ERROR:%v", err)
	}

	return resp.Result, nil
}
