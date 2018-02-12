package tgwrap

//
// GetUserProfilePhotosOpt represents optional params for GetUserProfilePhotos()
//
type GetUserProfilePhotosOpt struct {

	//
	// Sequential number of the first photo to be returned.
	// By default, all photos are returned.
	//
	Offset uint `json:"offset,omitempty"`

	//
	// Limits the number of photos to be retrieved.
	// Values between 1—100 are accepted. Defaults to 100.
	//
	Limit uint `json:"limit,omitempty"`
}

//
// GetUserProfilePhotos is used to get a list of profile pictures for a user.
// Returns a UserProfilePhotos object.
//
// userID: (int64) Unique identifier of the target user
//
// opt: (can be nil) optional params
//
func (p *bot) GetUserProfilePhotos(userID uint64, opt *GetUserProfilePhotosOpt) (*UserProfilePhotos, error) {

	type sendFormat struct {
		UserID                  uint64 `json:"user_id,omitempty"`
		GetUserProfilePhotosOpt `json:",omitempty"`
	}

	dataSend := sendFormat{
		UserID: userID,
	}

	if opt != nil {
		dataSend.GetUserProfilePhotosOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result *UserProfilePhotos `json:"result"`
	}

	err := p.getAPIResponse("getUserProfilePhotos", p.sendJSON, dataSend, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
