package tgwrap

// Contact represents a phone contact.
type Contact struct {

	// PhoneNumber is contact's phone number.
	PhoneNumber string `json:"phone_number"`

	// FirstName is contact's first name.
	FirstName string `json:"first_name"`

	// LastName. Optional. Contact's last name.
	LastName string `json:"last_name,omitempty"`

	// UserID. Optional. Contact's user identifier in Telegram.
	UserID uint64 `json:"user_id,omitempty"`
}
