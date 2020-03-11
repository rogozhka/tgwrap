package tgwrap

// ShippingAddress is a Telegram API object.
type ShippingAddress struct {

	//CountryCode - ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`

	//State (province) if applicable.
	State string `json:"state"`

	//City of the address.
	City string `json:"city"`

	//StreetLine1 - First line for the address.
	StreetLine1 string `json:"street_line1"`

	//StreetLine2 - Second line for the address.
	StreetLine2 string `json:"street_line2"`

	//PostCode - address post code.
	PostCode string `json:"post_code"`
}
