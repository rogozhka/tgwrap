package tgwrap

// User represents a Telegram user or bot.
type User struct {

	// User‘s or bot’s id
	ID int64 `json:"id"`

	// IsBot - true for bot.
	IsBot bool `json:"is_bot"`

	// FirstName. User‘s or bot’s first name.
	FirstName string `json:"first_name"`

	// LastName. Optional. User‘s or bot’s last name
	LastName string `json:"last_name"`

	// UserName. Optional. User‘s or bot’s username
	UserName string `json:"username,omitempty"`

	// LanguageCode. Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code,omitempty"`

	// CanJoinGroups. True, if the bot can be invited to groups. Returned only in getMe.
	CanJoinGroups bool `json:"can_join_groups,omitempty"`

	// CanReadAllGroupMessages. True, if privacy mode is disabled for the bot. Returned only in getMe.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`

	// SupportsInlineQueries. Optional. True, if the bot supports inline queries. Returned only in getMe.
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}
