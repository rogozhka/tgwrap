package tgwrap

// WebhookInfo contains information about the current status of a webhook.
type WebhookInfo struct {
	// URL may be empty if webhook is not set up.
	URL string `json:"url"`

	// HasCustomCertificate true, if a custom certificate was provided
	// for webhook certificate checks.
	HasCustomCertificate bool `json:"has_custom_certificate"`

	//PendingUpdateCount is number of updates awaiting delivery.
	PendingUpdateCount int `json:"pending_update_count"`

	//IPAddress is optional. Currently used webhook IP address.
	IPAddress string `json:"ip_address,omitempty"`

	// LastErrorDate is optional. Unix time for the most recent error
	// that happened when trying to deliver an update via webhook.
	LastErrorDate int `json:"last_error_date,omitempty"`

	// LastErrorMessage is optional. Error message in human-readable format
	// for the most recent error that happened when trying to deliver an update via webhook.
	LastErrorMessage string `json:"last_error_message,omitempty"`

	// LastSynchronizationErrorDate is optional. Unix time of the most recent error
	// that happened when trying to synchronize available updates with Telegram datacenters.
	LastSynchronizationErrorDate int `json:"last_synchronization_error_date,omitempty"`

	// MaxConnections is optional. Maximum allowed number of simultaneous
	// HTTPS connections to the webhook for update delivery.
	MaxConnections int `json:"max_connections,omitempty"`

	//AllowedUpdates is optional. A list of update types
	//the bot is subscribed to. Defaults to all update types except chat_member.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}
