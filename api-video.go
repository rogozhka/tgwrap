package tgwrap

// Video represents a video file.
type Video struct {

	// ID is unique identifier for this file.
	ID string `json:"file_id"`

	// Width. Video width as defined by sender.
	Width int `json:"width"`

	// Height. Video height as defined by sender.
	Height int `json:"height"`

	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`

	// Thumb. Optional. Video thumbnail.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// MIMEType. Optional. Mime type of a file as defined by sender
	MIMEType string `json:"mime_type,omitempty"`

	// FileSize. Optional. File size bytes.
	FileSize int `json:"file_size,omitempty"`
}
