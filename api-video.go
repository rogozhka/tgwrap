package tgwrap

//
// Telegram API object
//
// Video represents a video file.
//
type Video struct {

	//
	// Unique identifier for this file
	//
	ID string `json:"file_id"`

	//
	// Video width as defined by sender
	//
	Width uint `json:"width"`

	//
	// Video height as defined by sender
	//
	Height uint `json:"height"`

	//
	// Duration of the video in seconds as defined by sender
	//
	Duration uint `json:"duration"`

	//
	// Optional. Video thumbnail
	//
	Thumb *PhotoSize `json:"thumb,omitempty"`

	//
	// Optional. Mime type of a file as defined by sender
	//
	MIMEType string `json:"mime_type,omitempty"`

	//
	// Optional. File size
	//
	FileSize uint `json:"file_size,omitempty"`
}
