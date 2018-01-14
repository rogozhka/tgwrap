package tgwrap

type IBotMessages interface {
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
}
