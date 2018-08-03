package tgwrap

//
// PreCheckoutQuery contains information about an incoming pre-checkout query.
//
type PreCheckoutQuery struct {

    //
    // Unique query identifier
    //
    QueryID string `json:"id"`

    //
    // User who sent the query
    //
    From *User `json:"from"`

    //
    // Three-letter ISO 4217 currency code
    //
    Currency string `json:"currency"`

    //
    // Total price in the smallest units of the currency (integer, not float/double).
    //
    TotalAmount uint `json:"total_amount"`

    //
    // Total price in the smallest units of the currency (integer, not float/double).
    //
    InvoicePayload string `json:"invoice_payload"`

    //
    // Optional. Identifier of the shipping option chosen by the user
    //
    ShippingOptionID string `json:"shipping_option_id, omitempty"`

    //
    // Optional. Order info provided by the user
    //
    OrderInfo *OrderInfo `json:"order_info, omitempty"`
}
