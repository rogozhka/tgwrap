package tgwrap

type GetUserProfilePhotosOpt struct {
	commonRequestOptions

	// Offset is sequential number of the first photo to be returned.
	// By default, all photos are returned.
	Offset int `json:"offset,omitempty"`

	// Limit to limit the number of photos to be retrieved.
	// Values between 1—100 are accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`
}

// GetUserProfilePhotos to get a list of profile pictures for a user.
func (p *bot) GetUserProfilePhotos(userID int64, opt *GetUserProfilePhotosOpt) (*UserProfilePhotos, error) {

	type sendFormat struct {
		GetUserProfilePhotosOpt `json:",omitempty"`
		UserID                  int64 `json:"user_id"`
	}
	dataSend := sendFormat{
		UserID: userID,
	}
	if opt != nil {
		dataSend.GetUserProfilePhotosOpt = *opt
	} else {
		opt = &GetUserProfilePhotosOpt{}
	}
	var resp struct {
		GenericResponse
		Result *UserProfilePhotos `json:"result"`
	}
	err := p.getAPIResponse(opt.Context, "getUserProfilePhotos", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
