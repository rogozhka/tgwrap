package tgwrap

import (
	"fmt"
)

//
// RestrictChatMemberOpt represents optional params for RestrictChatMemberOpt()
//
type RestrictChatMemberOpt struct {

	//
	// Date when the user will be unbanned, unix time.
	// If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
	//
	UntilDate int64 `json:"until_date"`

	//
	// Pass True, if the user can send text messages, contacts, locations and venues
	//
	CanSendMessages bool `json:"can_send_messages"`

	//
	// Pass True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	//
	CanSendMediaMessages bool `json:"can_send_media_messages"`

	//
	// Pass True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
	//
	CanSendOtherMessages bool `json:"can_send_other_messages"`

	//
	// Pass True, if the user may add web page previews to their messages, implies can_send_media_messages
	//
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
}

// RestrictChatMember is used to restrict a user in a supergroup.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
// Pass True for all boolean parameters to lift restrictions from a user. Returns True on success.
//
// chatID: (uint64 or string) Unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// userID: (int64) Unique identifier of the target user
//
// opt: (can be nil) optional params
//
func (p *bot) RestrictChatMember(chatID interface{}, userID uint64, opt *RestrictChatMemberOpt) (bool, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		UserID uint64 `json:"user_id"`

		RestrictChatMemberOpt `json:",omitempty"`
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID),
		UserID: userID,
	}

	if opt != nil {
		dataSend.RestrictChatMemberOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result bool `json:"result"`
	}

	var sender fCommandSender = p.sendJSON

	err := p.getAPIResponse("restrictChatMember", sender, dataSend, &resp)
	if err != nil {
		return false, fmt.Errorf("getAPIResponse ERROR:%v", err)
	}

	return resp.Result, nil
}
