package tgwrap

//
// Telegram API object
//
type Chat struct {
	ID uint64 `json:"id"`

	//
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	//
	Type string `json:"type"`

	//
	// Optional. Title, for supergroups, channels and group chats
	//
	Title string `json:"title"`

	//
	// Optional. Username, for private chats, supergroups and channels if available
	//
	UserName string `json:"username"`

	//
	// Optional. First name of the other party in a private chat
	//
	FirstName string `json:"first_name"`

	//
	// Optional. Last name of the other party in a private chat
	//
	LastName string `json:"last_name"`

	//
	// Optional. True if a group has ‘All Members Are Admins’ enabled.
	//
	AllMembersAreAdministrators bool `json:"all_members_are_administrators"`

	//
	// Optional. Chat photo. Returned only in getChat.
	//
	Photo *ChatPhoto `json:"photo"`

	//
	// Optional. Description, for supergroups and channel chats. Returned only in getChat.
	//
	Description string `json:"description"`

	//
	// Optional. Chat invite link, for supergroups and channel chats. Returned only in getChat.
	//
	InviteLink string `json:"invite_link"`

	//
	// Optional. Pinned message, for supergroups and channel chats. Returned only in getChat.
	//
	PinnedMessage *Message `json:"pinned_message"`

	//
	// Optional. For supergroups, name of group sticker set. Returned only in getChat.
	//
	StickerSetName string `json:"sticker_set_name"`

	//
	// Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	//
	CanSetStickerSet bool `json:"can_set_sticker_set"`
}
