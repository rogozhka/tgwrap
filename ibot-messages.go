package tgwrap

//
// IBotMessages is just a group of methods
// to have smaller interface files
//
type IBotMessages interface {

	//
	// SendMessage is used method to send text messages.
	//
	// chatID: unique identifier for the target chat
	// or username(!) of the target channel (in the format @channelusername)
	//
	// text: text of the message to be sent
	//
	// opt: (can be nil) optional params
	//
	SendMessage(chatID interface{}, text string, opt *SendMessageOpt) (*Message, error)

	//
	// ForwardMessage is used to forward messages of any kind.
	// On success, the sent Message is returned.
	//
	// chatID: unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// fromChatID: unique identifier for the chat
	// where the original message was sent (or channel username
	// in the format @channelusername)
	//
	// disableNotification: sends the message silently.
	// Users will receive a notification with no sound.
	//
	// messageID: message identifier in the chat
	// specified in fromChatID
	//
	ForwardMessage(chatID interface{}, fromChatID interface{},
		disableNotification bool, messageID uint64) (*Message, error)
}
