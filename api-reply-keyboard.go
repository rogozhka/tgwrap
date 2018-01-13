package tgwrap

type ReplyKeyboardMarkup struct {
	Keyboard       [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard [][]*KeyboardButton `json:"resize_keyboard"`

	OneTimeKeyboard bool `json:"one_time_keyboard"`
	Selective       bool `json:"selective"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard [][]*InlineKeyboardButton `json:"remove_keyboard"`

	Selective bool `json:"selective"`
}

type KeyboardButton struct {
}
