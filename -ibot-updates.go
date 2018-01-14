package tgwrap

type IBotUpdates interface {
	//
	// Use GetUpdates method to receive incoming updates using long polling
	//
	// @param opt (optional, can be nil) GetUpdatesOpt with other optional params:
	// see struct description
	//
	GetUpdates(opt *GetUpdatesOpt) ([]*Update, error)
}
