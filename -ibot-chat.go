package tgwrap

type IBotChat interface {
	//
	// Use SendChatAction method when you need to tell the user
	// that something is happening on the bot's side. The status
	// is set for 5 seconds or less (when a message arrives from
	// your bot, Telegram clients clear its typing status).
	// Returns True on success.
	//
	// @param chatID Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// @param action Type of action to broadcast.
	//
	SendChatAction(chatID interface{}, action ChatActions) (bool, error)
}
