package tgwrap

//
// IBotContacts is a group of methods for sending contact-related
// information; used as a part of common IBot API interface
//
type IBotContacts interface {

	//
	// SendContact is used to to send phone contacts.
	// On success, the sent Message is returned.
	//
	// chatID: (uint64 or string) Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// phoneNumber: (string) Contact's phone number.
	//
	// firstName: (string) Contact's first name.
	//
	// opt: (can be nil) optional params
	//
	SendContact(chatID interface{}, phoneNumber string, firstName string, opt *SendContactOpt) (*Message, error)
}
