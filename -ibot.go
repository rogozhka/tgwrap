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

	IBotUpdates
	IBotMessages
	IBotChat
}
