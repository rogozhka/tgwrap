package tgwrap

// UserProfilePhotos represent a user's profile pictures.
type UserProfilePhotos struct {

	// TotalCount is total number of profile pictures the target user has.
	TotalCount int `json:"total_count"`

	// Photos - requested profile pictures (in up to 4 sizes each).
	Photos [][]PhotoSize `json:"photos"`
}
