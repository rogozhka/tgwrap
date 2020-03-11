package tgwrap

// Chat is a Telegram API object.
type Chat struct {

	// ID - unique identifier for this chat.
	ID int64 `json:"id"`

	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`

	// Title. Optional. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`

	// UserName. Optional. Username, for private chats, supergroups and channels if available
	UserName string `json:"username,omitempty"`

	// FirstName. Optional. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`

	// LastName. Optional. Last name of the other party in a private chat
	LastName string `json:"last_name,omitempty"`

	// Photo. Optional. Chat photo. Returned only in getChat.
	Photo *ChatPhoto `json:"photo,omitempty"`

	// Description. Optional. Description, for supergroups and channel chats. Returned only in getChat.
	Description string `json:"description,omitempty"`

	// InviteLink. Optional. Chat invite link, for supergroups and channel chats. Returned only in getChat.
	InviteLink string `json:"invite_link,omitempty"`

	// PinnedMessage. Optional. Pinned message, for supergroups and channel chats. Returned only in getChat.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Permissions. Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// SlowModeDelay. Optional. For supergroups, the minimum allowed delay between consecutive messages
	// sent by each unpriviledged user. Returned only in getChat.
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`

	// StickerSetName. Optional. For supergroups, name of group sticker set. Returned only in getChat.
	StickerSetName string `json:"sticker_set_name,omitempty"`

	// CanSetStickerSet. Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
}
