package tgwrap

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {

	//ID is unique query identifier.
	ID string `json:"id"`

	// From	is User who sent the query.
	From *User `json:"from"`

	//InvoicePayload - bot-specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	//ShippingAddress is user-specified shipping address.
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}
