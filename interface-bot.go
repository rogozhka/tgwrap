package tgwrap

//
// BotInterface is a summary declaration for nominal Telegram Bot API interface
//
type BotInterface interface {

	//
	// GetMe is used for testing your bot's auth token.
	// Requires no parameters. Returns basic information
	// about the bot in form of a User object.Telegram API method
	//
	GetMe() (*User, error)

	BotUpdatesInterface
	BotMessagesInterface
	BotChatInterface
	BotMediaInterface
	BotLocationInterface
	BotContactsInterface
}
