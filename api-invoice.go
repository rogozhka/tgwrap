package tgwrap

// Invoice contains basic information about an invoice.
type Invoice struct {

	// Title is a Product name.
	Title string `json:"title"`

	//
	// Description is a Product description
	//
	Description string `json:"description"`

	// StartParameter - Unique bot deep-linking parameter that can be used to generate this invoice.
	StartParameter string `json:"start_parameter"`

	// Currency - three-letter ISO 4217 [currency code](https://core.telegram.org/bots/payments#supported-currencies).
	Currency string `json:"currency"`

	// TotalAmount - total price in the smallest units of the currency
	// (integer, not float/double). For example, for a price
	// of US$ 1.45 pass amount = 145. See the exp parameter
	// in https://core.telegram.org/bots/payments/currencies.json
	// it shows the number of digits past the decimal point
	// for each currency (2 for the majority of currencies,
	// and zero for remaining ).
	TotalAmount int `json:"total_amount"`
}
