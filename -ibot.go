package tgwrap

//
// Telegram bot interface
//
type IBot interface {

	//
	// A simple method for testing your bot's auth token.
	// Requires no parameters. Returns basic information about the bot
	//
	GetMe() (*User, error)

	//
	// Use GetUpdates method to receive incoming updates using long polling
	//
	// @param opt (optional, can be nil) GetUpdatesOpt with other optional params:
	// see struct description
	//
	GetUpdates(opt *GetUpdatesOpt) ([]*Update, error)

	//
	// Use SendMessage method to send text messages.
	//
	// @param chatID Unique identifier for the target chat
	// or username(!) of the target channel (in the format @channelusername)
	//
	// @param text Text of the message to be sent
	//
	// @param opt (optional, can be nil) SendMessageOpt with other optional params:
	// see struct description
	//
	SendMessage(chatID interface{}, text string, opt *SendMessageOpt) (*Message, error)

	//
	// Use ForwardMessage to forward messages of any kind.
	// On success, the sent Message is returned.
	//
	// @param chatID Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// @param fromChatID Unique identifier for the chat
	// where the original message was sent (or channel username
	// in the format @channelusername)
	//
	// @param disableNotification Sends the message silently.
	// Users will receive a notification with no sound.
	//
	// @param messageID Message identifier in the chat
	// specified in fromChatID
	//
	ForwardMessage(chatID interface{}, fromChatID interface{},
		disableNotification bool, messageID uint64) (*Message, error)

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
