package tgwrap

//
// Telegram API object
//
// This object represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
//
type InlineQuery struct {

	//
	// Unique identifier for this query
	//
	ID string `json:"id"`

	//
	// Sender
	//
	From *User `json:"from"`

	//
	// Optional. Sender location, only for bots that request user location
	//
	Location *Location `json:"location,omitempty"`

	//
	// Text of the query (up to 512 characters)
	//
	Query string `json:"query"`

	//
	// Offset of the results to be returned, can be controlled by the bot
	//
	Offset string `json:"offset"`
}
