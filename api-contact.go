package tgwrap

//
// Contact represents a phone contact.
//
type Contact struct {

	//
	// Contact's phone number
	//
	PhoneNumber string `json:"phone_number"`

	//
	// Contact's first name
	//
	FirstName string `json:"first_name"`

	//
	// Optional. Contact's last name
	//
	LastName string `json:"last_name,omitempty"`

	//
	// Optional. Contact's user identifier in Telegram
	//
	UserID uint64 `json:"user_id,omitempty"`
}
