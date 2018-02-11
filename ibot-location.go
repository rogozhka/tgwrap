package tgwrap

//
// IBotLocation is a group of methods for location
// used as a part of common IBot API interface
//
type IBotLocation interface {

	// SendLocation is used to send point on the map.
	// On success, the sent Message is returned.
	//
	// chatID: (uint64 or string) Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// latitude: (float64) Latitude of the location.
	//
	// longitude: (float64) Longitude of the location.
	//
	// opt: (can be nil) optional params
	//
	SendLocation(chatID interface{}, latitude float64, longitude float64, opt *SendLocationOpt) (*Message, error)

	// SendVenue is used to send information about a venue.
	// On success, the sent Message is returned.
	//
	// chatID: (uint64 or string) Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// latitude: (float64) Latitude of the venue.
	//
	// longitude: (float64) Longitude of the venue.
	//
	// title: (string) Name of the venue.
	//
	// address: (string) Address of the venue.
	//
	// opt: (can be nil) optional params
	//
	SendVenue(chatID interface{}, latitude float64, longitude float64, title string, address string, opt *SendVenueOpt) (*Message, error)

	//
	// StopMessageLiveLocation is used to stop updating a live location message sent by the bot or via the bot (for inline bots) before live_period expires.
	// On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
	//
	// opt: (can be nil) optional params
	//
	StopMessageLiveLocation(opt *StopMessageLiveLocationOpt) (interface{}, error)
}
