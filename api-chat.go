package tgwrap

//
// Chat is a Telegram API object
//
type Chat struct {

	//
	// Unique identifier for this chat.
	//
	ID int64 `json:"id"`

	//
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	//
	Type string `json:"type"`

	//
	// Optional. Title, for supergroups, channels and group chats
	//
	Title string `json:"title,omitempty"`

	//
	// Optional. Username, for private chats, supergroups and channels if available
	//
	UserName string `json:"username,omitempty"`

	//
	// Optional. First name of the other party in a private chat
	//
	FirstName string `json:"first_name,omitempty"`

	//
	// Optional. Last name of the other party in a private chat
	//
	LastName string `json:"last_name,omitempty"`

	//
	// Optional. True if a group has ‘All Members Are Admins’ enabled.
	//
	AllMembersAreAdministrators bool `json:"all_members_are_administrators,omitempty"`

	//
	// Optional. Chat photo. Returned only in getChat.
	//
	Photo *ChatPhoto `json:"photo,omitempty"`

	//
	// Optional. Description, for supergroups and channel chats. Returned only in getChat.
	//
	Description string `json:"description,omitempty"`

	//
	// Optional. Chat invite link, for supergroups and channel chats. Returned only in getChat.
	//
	InviteLink string `json:"invite_link,omitempty"`

	//
	// Optional. Pinned message, for supergroups and channel chats. Returned only in getChat.
	//
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	//
	// Optional. For supergroups, name of group sticker set. Returned only in getChat.
	//
	StickerSetName string `json:"sticker_set_name,omitempty"`

	//
	// Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	//
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
}
