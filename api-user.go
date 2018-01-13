package tgwrap

//
// Telegram API object
//
// User represents a Telegram user or bot.
//
type User struct {
	//
	// User‘s or bot’s id
	//
	ID uint64 `json:"id"`

	//
	// True for bot
	//
	IsBot bool `json:"is_bot"`

	//
	// User‘s or bot’s first name
	//
	FirstName string `json:"first_name"`

	//
	// Optional. User‘s or bot’s last name
	//
	LastName string `json:"last_name"`

	//
	// Optional. User‘s or bot’s username
	//
	UserName string `json:"username"`

	//
	// Optional. IETF language tag of the user's language
	//
	LanguageCode string `json:"language_code"`
}
