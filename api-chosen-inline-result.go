package tgwrap

// ChosenInlineResult represents a result of an inline query that was chosen
// by the user and sent to their chat partner.
type ChosenInlineResult struct {

	// ResultID - unique identifier for the result that was chosen.
	ResultID string `json:"result_id"`

	// From is the user that chose the result.
	From *User `json:"from"`

	// Location. Optional. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`

	// InlineMessageID. Optional. Identifier of the sent inline message.
	// Available only if there is an inline keyboard attached to the message.
	// Will be also received in callback queries and can be used to edit the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Query that was used to obtain the result.
	Query string `json:"query"`
}
