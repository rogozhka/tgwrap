package tgwrap

import (
	"fmt"
)

//
// SendVenueOpt represents optional params for SendVenue
//
type SendVenueOpt struct {

	//
	// Foursquare identifier of the venue.
	//
	FoursquareID uint `json:"foursquare_id,omitempty"`

	//
	// Sends the message silently. Users will receive a notification with no sound.
	//
	DisableNotification bool `json:"disable_notification,omitempty"`

	//
	// If the message is a reply, ID of the original message
	//
	ReplyToID uint64 `json:"reply_to_message_id,omitempty"`

	//
	// Additional interface options. A JSON-serialized object
	// for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard
	// or to force a reply from the user.
	//
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

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
func (p *bot) SendVenue(chatID interface{}, latitude float64, longitude float64, title string, address string, opt *SendVenueOpt) (*Message, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		SendVenueOpt `json:",omitempty"`

		Latitude float64 `json:"latitude"`

		Longitude float64 `json:"longitude"`

		Title string `json:"title"`

		Address string `json:"address"`
	}

	dataSend := sendFormat{
		ChatID:    fmt.Sprint(chatID),
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}

	if opt != nil {
		dataSend.SendVenueOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result *Message `json:"result"`
	}

	var sender fCommandSender = p.sendJSON

	err := p.getAPIResponse("sendVenue", sender, dataSend, &resp)
	if err != nil {
		return nil, fmt.Errorf("getAPIResponse ERROR:%v", err)
	}

	return resp.Result, nil
}