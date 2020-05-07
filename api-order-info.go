package tgwrap

// OrderInfo represents information about an order.
type OrderInfo struct {

	//Name.	Optional. User name.
	Name string `json:"name,omitempty"`

	//PhoneNumber. Optional. User's phone number.
	PhoneNumber string `json:"phone_number,omitempty"`

	//Email. Optional. User email.
	Email string `json:"email,omitempty"`

	//ShippingAddress. Optional. User shipping address.
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}
