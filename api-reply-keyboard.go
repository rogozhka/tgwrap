package tgwrap

//
// Telegram API object
//
// ReplyKeyboardMarkup represents a custom keyboard with reply options
// (see Introduction to bots for details and examples).
//
type ReplyKeyboardMarkup struct {

	//
	// Array of button rows, each represented by an Array of KeyboardButton objects
	//
	Keyboard [][]*KeyboardButton `json:"keyboard"`

	//
	// Optional. Requests clients to resize the keyboard vertically for optimal fit
	// (e.g., make the keyboard smaller if there are just two rows of buttons).
	// Defaults to false, in which case the custom keyboard is always
	// of the same height as the app's standard keyboard.
	//
	ResizeKeyboard [][]*KeyboardButton `json:"resize_keyboard,omitempty"`

	//
	// Optional. Requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display
	// the usual letter-keyboard in the chat – the user can press a special button
	// in the input field to see the custom keyboard again. Defaults to false.
	//
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	//
	// Optional. Use this parameter if you want to show the keyboard to specific users only.
	// Targets:
	// 1) users that are @mentioned in the text of the Message object;
	//
	// 2) if the bot's message is a reply (has reply_to_message_id),
	// sender of the original message.
	//
	Selective bool `json:"selective,omitempty"`
}

//
// Telegram API object
//
// Upon receiving a message with this object, Telegram clients will
// remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after
// the user presses a button (see ReplyKeyboardMarkup).
//
type ReplyKeyboardRemove struct {
	RemoveKeyboard [][]*InlineKeyboardButton `json:"remove_keyboard"`

	Selective bool `json:"selective"`
}

//
// This object represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object
// to specify text of the button. Optional fields are mutually exclusive.
//
// Note: request_contact and request_location options will only work
// in Telegram versions released after 9 April, 2016. Older clients will ignore them.
//
type KeyboardButton struct {

	//
	// Text of the button. If none of the optional fields are used,
	// it will be sent as a message when the button is pressed
	//
	Text string `json:"text"`

	//
	// Optional. If True, the user's phone number will be sent
	// as a contact when the button is pressed. Available in private chats only
	//
	RequestContact string `json:"request_contact,omitempty"`

	//
	// Optional. If True, the user's current location will be sent
	// when the button is pressed. Available in private chats only
	//
	RequestLocation string `json:"request_location,omitempty"`
}
