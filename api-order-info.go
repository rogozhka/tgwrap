package tgwrap

//
// OrderInfo represents information about an order.
//
type OrderInfo struct {

    //
    // Optional. User name
    //
    Name string `json:"name,omitempty"`

    //
    // Optional. User's phone number
    //
    PhoneNumber string `json:"phone_number,omitempty"`

    //
    // Optional. User email
    //
    Email string `json:"email,omitempty"`

    //
    // Optional. User shipping address
    //
    ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}
