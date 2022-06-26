package tgwrap

import (
	"context"
	"fmt"
	"reflect"
)

// SendLocationOpt represents optional params for SendLocation.
type SendLocationOpt struct {
	commonRequestOptions

	// LivePeriod is period in seconds for which the location
	// will be updated (see Live Locations), should be between 60 and 86400.
	LivePeriod uint `json:"live_period,omitempty"`

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

// SendLocation is used to send point on the map.
// On success, the sent Message is returned.
//
// chatID: (int64 or string) is unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// latitude: (float64) Latitude of the location.
//
// longitude: (float64) Longitude of the location.
//
// opt: (can be nil) optional params
//
func (p *bot) SendLocation(chatID interface{}, latitude float64, longitude float64, opt *SendLocationOpt) (*Message, error) {

	type sendFormat struct {
		ChatID          string `json:"chat_id"`
		SendLocationOpt `json:",omitempty"`
		Latitude        float64 `json:"latitude"`
		Longitude       float64 `json:"longitude"`
	}

	dataSend := sendFormat{
		ChatID:    fmt.Sprint(chatID),
		Latitude:  latitude,
		Longitude: longitude,
	}

	if opt == nil {
		opt = &SendLocationOpt{}
	}
	if opt.Context == nil {
		opt.Context = context.Background()
	}

	if lp := reflect.ValueOf(opt.LivePeriod); lp.IsValid() {
		if lp.Uint() < 60 && lp.Uint() > 86400 {
			return nil, fmt.Errorf("invalid live period")
		}
	}
	dataSend.SendLocationOpt = *opt
	var resp struct {
		GenericResponse
		Result *Message `json:"result"`
	}
	err := p.getAPIResponse(opt.Context, "sendLocation", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
