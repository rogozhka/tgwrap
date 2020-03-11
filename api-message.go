package tgwrap

import "time"

// Message is a Telegram API object.
type Message struct {

	// ID is unique message identifier inside chat.
	ID int64 `json:"message_id"`

	// From. Optional. Sender, empty for messages sent to channels.
	From *User `json:"from,omitempty"`

	// Date the message was sent in Unix time.
	Date int64 `json:"date"`

	// Chat - conversation the message belongs to.
	Chat *Chat `json:"chat"`

	// ForwardFrom. Optional. For forwarded messages, sender of the original message.
	ForwardFrom *User `json:"forward_from,omitempty"`

	// ForwardFromChat. Optional. For messages forwarded from channels, information about the original channel.
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`

	// ForwardFromMessageID. Optional. For messages forwarded from channels,
	// identifier of the original message in the channel.
	ForwardFromMessageID int64 `json:"forward_from_message_id,omitempty"`

	// ForwardSignature. Optional. For messages forwarded from channels, signature of the post author if present.
	ForwardSignature string `json:"forward_signature,omitempty"`

	// ForwardSignature. Optional. Optional. Sender's name for messages forwarded from users
	// who disallow adding a link to their account in forwarded messages.
	ForwardSenderName string `json:"forward_sender_name,omitempty"`

	// ForwardDate. Optional. For forwarded messages, date the original message was sent in Unix time.
	ForwardDate int64 `json:"forward_date,omitempty"`

	// ReplyToMessage. Optional. For replies, the original message.
	// Note that the Message object in this field
	// will not contain further reply_to_message fields
	// even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// EditDate. Optional. Date the message was last edited in Unix time.
	EditDate int64 `json:"edit_date,omitempty"`

	// MediaGroupID. Optional. The unique identifier of a media message group this message belongs to.
	MediaGroupID string `json:"media_group_id,omitempty"`

	// AuthorSignature. Optional. Signature of the post author for messages in channels.
	AuthorSignature string `json:"author_signature,omitempty"`

	// Text. Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Text string `json:"text,omitempty"`

	// Entities. Optional. For text messages, special entities
	// like usernames, URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`

	// CaptionEntities. Optional. For messages with a caption, special entities
	// like usernames, URLs, bot commands, etc. that appear in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Audio. Optional. Message is an audio file, information about the file.
	Audio *Audio `json:"audio,omitempty"`

	// Document. Optional. Message is a general file, information about the file.
	Document *Document `json:"document,omitempty"`

	// Animation. Optional. Message is an animation, information about the animation.
	// For backward compatibility, when this field is set, the document field will also be set.
	Animation *Animation `json:"animation,omitempty"`

	// Game. Optional. Message is a game, information about the game.
	Game *Game `json:"game,omitempty"`

	// Photo. Optional. Message is a photo, available sizes of the photo.
	Photo []PhotoSize `json:"photo,omitempty"`

	// Sticker. Optional. Message is a sticker, information about the sticker.
	Sticker *Sticker `json:"sticker,omitempty"`

	// Video. Optional. Message is a video, information about the video.
	Video *Video `json:"video,omitempty"`

	// Voice. Optional. Message is a voice message, information about the file.
	Voice *Voice `json:"voice,omitempty"`

	// VideoNote. Optional. Message is a video note, information about the video message.
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Caption. Optional. Caption for the audio, document, photo, video or voice, 0-200 characters.
	Caption string `json:"caption,omitempty"`

	// Contact. Optional. Message is a shared contact, information about the contact.
	Contact *Contact `json:"contact,omitempty"`

	// Location. Optional. Message is a shared location, information about the location.
	Location *Location `json:"location,omitempty"`

	// Venue. Optional. Message is a venue, information about the venue.
	Venue *Venue `json:"venue,omitempty"`

	// Poll. Optional. Message is a native poll, information about the poll.
	Poll *Poll `json:"poll,omitempty"`

	// NewChatMembers. Optional. New members that were added to the group
	// or supergroup and information about them
	// (the bot itself may be one of these members).
	NewChatMembers []User `json:"new_chat_members,omitempty"`

	// LeftChatMembers. Optional. A member was removed from the group,
	// information about them (this member may be the bot itself).
	LeftChatMembers *User `json:"left_chat_member,omitempty"`

	// NewChatTitle. Optional. A chat title was changed to this value.
	NewChatTitle string `json:"new_chat_title,omitempty"`

	// NewChatPhoto. Optional. A chat photo was changed to this value.
	NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`

	// DeleteChatPhoto. Optional. Service message: the chat photo was deleted.
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`

	// GroupChatCreated. Optional. Service message: the group has been created.
	GroupChatCreated bool `json:"group_chat_created,omitempty"`

	// SupergroupChatCreated. Optional. Service message: the supergroup has been created.
	// This field can‘t be received in a message coming through updates,
	// because bot can’t be a member of a supergroup when it is created.
	// It can only be found in reply_to_message if someone replies
	// to a very first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`

	// ChannelChatCreated. Optional. Service message: the channel has been created.
	// This field can‘t be received in a message coming through updates,
	// because bot can’t be a member of a channel when it is created.
	// It can only be found in reply_to_message if someone replies
	// to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`

	// MigrateChatID. Optional. The group has been migrated to a supergroup
	// with the specified identifier. This number may be greater
	// than 32 bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But
	// it is smaller than 52 bits, so a signed 64 bit
	// integer or double-precision float type are safe
	// for storing this identifier.
	MigrateChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// MigrateFromChatID. Optional. The supergroup has been migrated from a group
	// with the specified identifier. This number may be greater
	// than 32 bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But
	// it is smaller than 52 bits, so a signed 64 bit
	// integer or double-precision float type are safe
	// for storing this identifier.
	MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`

	// PinnedMessage. Optional. Specified message was pinned.
	// Note that the Message object in this field
	// will not contain further reply_to_message fields
	// even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Invoice. Optional. Message is an invoice for a payment,
	// information about the invoice.
	Invoice *Invoice `json:"invoice,omitempty"`

	// SuccessfulPayment. Optional. Message is a service message about
	// a successful payment, information about the payment.
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`

	// ConnectedWebsite. The domain name of the website on which the user has logged in.
	ConnectedWebsite string `json:"connected_website,omitempty"`

	//PassportData. Optional. Telegram Passport data.
	PassportData *PassportData `json:"passport_data,omitempty"`

	// ReplyMarkup. Optional. Inline keyboard attached to the message.
	// login_url buttons are represented as ordinary url buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Time returns Message.Date in time.Time format
func (p *Message) Time() time.Time {
	return time.Unix(p.Date, 0)
}

// ForwardTime returns Message.ForwardDate in time.Time format.
func (p *Message) ForwardTime() time.Time {
	return time.Unix(p.ForwardDate, 0)
}

// EditTime returns Message.EditDate in time.Time format.
func (p *Message) EditTime() time.Time {
	return time.Unix(p.EditDate, 0)
}
