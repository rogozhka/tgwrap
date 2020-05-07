package tgwrap

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {

	// CanSendMessages.	Optional. True, if the user is allowed to send text messages, contacts, locations and venues.
	CanSendMessages bool `json:"can_send_messages"`

	// CanSendMediaMessages.	Optional. True, if the user is allowed to send audios, documents, photos,
	// videos, video notes and voice notes, implies can_send_messages.
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

	// CanSendPolls.Optional. True, if the user is allowed to send polls, implies can_send_messages.
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// CanSendOtherMessages. Optional. True, if the user is allowed to send animations,
	// games, stickers and use inline bots, implies can_send_media_messages.
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// CanAddWebPagePreviews. Optional. True, if the user is allowed to add web page previews
	// to their messages, implies can_send_media_messages.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// CanChangeInfo. Optional. True, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups.
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// CanInviteUsers. Optional. True, if the user is allowed to invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// CanPinMessages. Optional. True, if the user is allowed to pin messages. Ignored in public supergroups.
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}
