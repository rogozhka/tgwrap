package tgwrap

import (
	"context"
	"fmt"
)

// SendVenueOpt represents optional params for SendVenue.
type SendVenueOpt struct {
	commonRequestOptions

	// Foursquare is identifier of the venue.
	FoursquareID string `json:"foursquare_id,omitempty"`

	// DisableNotification sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// ReplyToID is optional ID of the original message if the message is a reply.
	ReplyToID int64 `json:"reply_to_message_id,omitempty"`

	// ReplyMarkup - additional interface options. A JSON-serialized object
	// for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard
	// or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// SendVenue is used to send information about a venue.
// On success, the sent Message is returned.
//
// chatID: (int64 or string) is unique identifier for the target chat
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
func (p *bot) SendVenue(chatID interface{},
	latitude float64,
	longitude float64,
	title string,
	address string,
	opt *SendVenueOpt,
) (*Message, error) {

	type sendFormat struct {
		ChatID       string `json:"chat_id"`
		SendVenueOpt `json:",omitempty"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
		Title        string  `json:"title"`
		Address      string  `json:"address"`
	}
	dataSend := sendFormat{
		ChatID:    fmt.Sprint(chatID),
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}

	if opt == nil {
		opt = &SendVenueOpt{}
	}
	if opt.Context == nil {
		opt.Context = context.Background()
	}
	dataSend.SendVenueOpt = *opt

	var resp struct {
		GenericResponse
		Result *Message `json:"result"`
	}
	err := p.getAPIResponse(opt.Context, "sendVenue", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
