package tgwrap

//
// Telegram API object
//
type MessageEntity struct {
	//
	// Type of the entity. Can be
	// mention (@username),
	// hashtag,
	// bot_command,
	// url,
	// email,
	// bold (bold text),
	// italic (italic text),
	// code (monowidth string),
	// pre (monowidth block),
	// text_link (for clickable text URLs),
	// text_mention (for users without usernames)
	//
	Type string `json:"type"`

	//
	// Offset in UTF-16 code units to the start of the entity
	//
	Offset int `json:"offset"`

	//
	// Length of the entity in UTF-16 code units
	//
	Length int `json:"length"`

	//
	// Optional. For “text_link” only, url that will be opened after user taps on the text
	//
	URL string `json:"url"`

	//
	// Optional. For “text_mention” only, the mentioned user
	//
	User *User `json:"user"`
}
