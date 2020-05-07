package tgwrap

// Location represents a point on the map.
type Location struct {

	// Longitude as defined by sender.
	Longitude float64 `json:"longitude"`

	// Latitude as defined by sender.
	Latitude float64 `json:"latitude"`
}
