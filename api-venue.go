package tgwrap

// Venue represents a place, meeting point, event point.
type Venue struct {

	// Location is Venue location.
	Location *Location `json:"location"`

	// Title is Name of the venue.
	Title string `json:"title"`

	// Address of the venue.
	Address string `json:"address"`

	// FoursquareID. Optional. Foursquare identifier of the venue.
	FoursquareID string `json:"foursquare_id,omitempty"`

	// FoursquareType. Optional. Foursquare type of the venue.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type"`
}
