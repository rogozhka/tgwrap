package tgwrap

//
// IBotChat is a group of methods for chat manipulation
// used as a part of common IBot API interface
//
type IBotChat interface {

	//
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
	//
	SendChatAction(chatID interface{}, action ChatActions) (bool, error)

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
	RestrictChatMember(chatID interface{}, userID uint64, opt *RestrictChatMemberOpt) (bool, error)
}
