package tgwrap

//
// IBotUpdates is a group of methods related to getting updates
// used as a part of common IBot API interface
//
type IBotUpdates interface {

	//
	// GetUpdates is used to receive incoming updates using [long] polling
	//
	// opt: (can be nil) optional params
	//
	GetUpdates(opt *GetUpdatesOpt) ([]*Update, error)
}
