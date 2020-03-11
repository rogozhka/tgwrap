package tgwrap

// MessageEntity represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {

	// Type of the entity.
	// Can be:
	// - mention (@username),
	// - hashtag (#hashtag),
	// - cashtag ($USD),
	// - bot_command (/start@jobs_bot),
	// - url,
	// - email,
	// - bold (bold text),
	// - italic (italic text),
	// - underline (underline text),
	// - strikethrough (strikethrough text),
	// - code (monowidth string),
	// - pre (monowidth block),
	// - text_link (for clickable text URLs),
	// - text_mention (for users without usernames).
	Type string `json:"type"`

	// Offset in UTF-16 code units to the start of the entity.
	Offset int `json:"offset"`

	// Length of the entity in UTF-16 code units.
	Length int `json:"length"`

	// URL. Optional. For “text_link” only, url that will be opened after user taps on the text.
	URL string `json:"url,omitempty"`

	// User. Optional. For “text_mention” only, the mentioned user.
	User *User `json:"user,omitempty"`

	// Language is Optional. For “pre” only, the programming language of the entity text.
	Language string `json:"language,omitempty"`
}
