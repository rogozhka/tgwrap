package tgwrap

type UpdateType string

const (
	// UpdateTypeMessage is new incoming message of any kind — text, photo, sticker, etc.
	UpdateTypeMessage = "message"

	// UpdateTypeEditedMessage is new version of a message that is known to the bot and was edited.
	UpdateTypeEditedMessage = "edited_message"

	// UpdateTypeChannelPost is new incoming channel post of any kind — text, photo, sticker, etc.
	UpdateTypeChannelPost = "channel_post"

	// UpdateTypeEditedChannelPost is new version of a channel post that is known to the bot and was edited.
	UpdateTypeEditedChannelPost = "edited_channel_post"

	// UpdateTypeinlineQuery is new incoming inline query.
	UpdateTypeinlineQuery = "inline_query"

	// UpdateTypeChosenInlineResult - the result of an inline query that was chosen
	// by a user and sent to their chat partner.
	UpdateTypeChosenInlineResult = "chosen_inline_result"

	// UpdateTypeCallbackQuery is new incoming callback query/.
	UpdateTypeCallbackQuery = "callback_query"

	// UpdateTypeShippingQuery
	UpdateTypeShippingQuery = "shipping_query"

	// UpdateTypePreCheckoutQuery is new incoming shipping query.
	UpdateTypePreCheckoutQuery = "pre_checkout_query"

	// UpdateTypePoll is new poll state. Bots receive only updates
	// about stopped polls and polls, which are sent by the bot.
	UpdateTypePoll = "poll"

	// UpdateTypePollAnswer - a user changed their answer in a non-anonymous poll.
	// Bots receive new votes only in polls that were sent by the bot itself
	UpdateTypePollAnswer = "poll_answer"
)

// Update represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {

	// ID is the update‘s unique identifier. Update identifiers start
	// from a certain positive number and increase sequentially. This ID
	// becomes especially handy if you’re using Webhooks, since it allows
	// you to ignore repeated updates or to restore the correct update sequence,
	// should they get out of order. If there are no new updates for at least a week,
	// then identifier of the next update will be chosen randomly instead of sequentially.
	ID uint64 `json:"update_id"`

	// Message. Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`

	// EditedMessage. Optional. New version of a message that is known to the bot and was edited.
	EditedMessage *Message `json:"edited_message,omitempty"`

	// ChannelPost. Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`

	// EditedChannelPost. Optional. New version of a channel post that is known to the bot and was edited.
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// InlineQuery. Optional. New incoming inline query.
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`

	// ChosenInlineResult. Optional. The result of an inline query that was chosen
	// by a user and sent to their chat partner. Please see our
	// documentation on the feedback collecting for details
	// on how to enable these updates for your bot.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

	// CallbackQuery. Optional. New incoming callback query.
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

	// ShippingQuery. Optional. New incoming shipping query.
	// Only for invoices with flexible price.
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	// PreCheckoutQuery. Optional. New incoming pre-checkout query.
	// Contains full information about checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	// Poll. Optional. New poll state. Bots receive only updates
	// about stopped polls and polls, which are sent by the bot.
	Poll *Poll `json:"poll,omitempty"`

	// PollAnswer. Optional. A user changed their answer in a non-anonymous poll.
	// Bots receive new votes only in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
}
