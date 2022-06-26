package tgwrap

// BotUpdatesInterface is a group of methods related to getting updates
// used as a part of common BotInterface API interface
type BotUpdatesInterface interface {

	// GetUpdates is used to receive incoming updates using [long] polling
	//
	// opt: (can be nil) optional params
	//
	GetUpdates(opt *GetUpdatesOpt) ([]Update, error)
}
