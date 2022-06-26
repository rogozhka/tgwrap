package tgwrap

import (
	"context"
	"reflect"

	"github.com/rogozhka/thestruct"
)

// SetWebhookOpt represents optional params for SetWebhook.
type SetWebhookOpt struct {
	commonRequestOptions

	// Certificate is optional file, your public key certificate
	// so that the root certificate in use can be checked.
	// See self-signed guide for details.
	// https://core.telegram.org/bots/self-signed
	Certificate *InputFile `json:"certificate,omitempty"`

	// IPAddress is optional. The fixed IP address
	// which will be used to send webhook requests
	// instead of the IP address resolved through DNS.
	IPAddress string `json:"ip_address,omitempty"`

	// MaxConnections is maximum allowed number of simultaneous HTTPS connections
	// to the webhook for update delivery, 1-100. Defaults to 40.
	// Use lower values to limit the load on your bot‘s server,
	// and higher values to increase your bot’s throughput.
	MaxConnections int `json:"max_connections,omitempty"`

	// AllowedUpdates lists types of updates you want your bot to receive.
	// For example, specify [“message”, “edited_channel_post”,
	// “callback_query”] to only receive updates of these types.
	// See UpdateType for a complete list of available update types.
	// Specify an empty list to receive all updates regardless
	// of type (default). If not specified, the previous setting
	// will be used.
	//
	// Please note that this parameter doesn't affect updates
	// created before the call to the setWebhook, so unwanted
	// updates may be received for a short period of time.
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`

	// DropPendingUpdates is optional. Pass True to drop all pending updates.
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// SetWebhook is used to specify a url and receive incoming updates
// via an outgoing webhook. Whenever there is an update for the bot,
// we will send an HTTPS POST request to the specified url, containing
// a JSON-serialized Update. In case of an unsuccessful request,
// we will give up after a reasonable amount of attempts. Returns true.
//
// If you'd like to make sure that the Webhook request comes from Telegram,
// we recommend using a secret path in the URL, e.g. https://www.example.com/<token>.
// Since nobody else knows your bot‘s token, you can be pretty sure it’s us.
//
// Use an empty string as url to remove webhook integration.
func (p *bot) SetWebhook(url string, opt *SetWebhookOpt) (bool, error) {

	type sendFormat struct {
		URL           string `json:"url"`
		SetWebhookOpt        // optional part
	}
	dataSend := sendFormat{
		URL: url,
	}
	resp := struct {
		GenericResponse
		Result bool `json:"result"`
	}{}

	if opt == nil {
		opt = &SetWebhookOpt{}
	}
	if opt.Context == nil {
		opt.Context = context.Background()
	}
	dataSend.SetWebhookOpt = *opt

	sender := p.sendJSON
	tt := thestruct.Type(reflect.TypeOf(dataSend.Certificate))
	if "InputFileLocal" == tt.Name() {
		sender = p.sendFormData
	}
	err := p.getAPIResponse(opt.Context, "setWebhook", sender, dataSend, &resp)
	return resp.Result, err
}
