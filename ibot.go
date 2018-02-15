package tgwrap

//
// IBot is a declaration for nominal Telegram Bot API interface
//
type IBot interface {

	//
	// GetMe is used for testing your bot's auth token.
	// Requires no parameters. Returns basic information
	// about the bot in form of a User object.Telegram API method
	//
	GetMe() (*User, error)

	IBotUpdates
	IBotMessages
	IBotChat
	IBotMedia
	IBotLocation
	IBotContacts
}
