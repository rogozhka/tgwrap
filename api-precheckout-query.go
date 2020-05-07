package tgwrap

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {

	//ID is unique query identifier.
	ID string `json:"id"`

	//From is user who sent the query.
	From *User `json:"from"`

	//Currency is a three-letter ISO 4217 currency code.
	Currency string `json:"currency"`

	// TotalAmount - total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json,
	// it shows the number of digits past the decimal point
	// for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	//InvoicePayload - bot-specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	//ShippingOptionID. Optional. Identifier of the shipping option chosen by the user.
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	//OrderInfo. Optional. Order info provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}
