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

	// UnbanChatMember is used to unban a previously kicked user in a supergroup or channel.
	// The user will not return to the group or channel automatically, but will be able to join via link, etc.
	// The bot must be an administrator for this to work.
	// Returns True on success.
	//
	// chatID: (uint64 or string) Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// userID: (int64) Unique identifier of the target user
	//
	UnbanChatMember(chatID interface{}, userID uint64) (bool, error)
}
