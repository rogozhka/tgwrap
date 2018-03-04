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
	PromoteChatMember(chatID interface{}, userID uint64, opt *PromoteChatMemberOpt) (bool, error)
}
