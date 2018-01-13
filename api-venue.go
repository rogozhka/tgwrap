package tgwrap

//
// Telegram API object
//
// Venue represents a place, meeting point, event point
//
type Venue struct {

	//
	// Venue location
	//
	Location *Location `json:"location"`

	//
	// Name of the venue
	//
	Title string `json:"title"`

	//
	// Address of the venue
	//
	Address string `json:"address"`

	//
	// Optional. Foursquare identifier of the venue
	//
	FoursquareID string `json:"foursquare_id,omitempty"`
}
