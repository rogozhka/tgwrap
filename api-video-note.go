package tgwrap

//
// Telegram API object
//
// VideoNote represents a video message (available in Telegram apps as of v.4.0).
//
type VideoNote struct {

	//
	// Unique identifier for this file
	//
	ID string `json:"file_id"`

	//
	// Video width and height as defined by sender
	//
	Length uint `json:"length"`

	//
	// Duration of the video in seconds as defined by sender
	//
	Duration uint `json:"duration"`

	//
	// Optional. Video thumbnail
	//
	Thumb *PhotoSize `json:"thumb,omitempty"`

	//
	// Optional. File size
	//
	FileSize uint `json:"file_size,omitempty"`
}
