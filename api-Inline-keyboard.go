package tgwrap

//
// This object represents an inline keyboard that appears
// right next to the message it belongs to.
//
type InlineKeyboardMarkup struct {

	//
	// Array of button rows, each represented
	// by an Array of InlineKeyboardButton objects
	//
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

//
//
//
type InlineKeyboardButton struct {
}
