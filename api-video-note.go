package tgwrap

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {

	// ID is unique identifier for this file.
	ID string `json:"file_id"`

	// Length. Video width and height as defined by sender
	Length int `json:"length"`

	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`

	// Thumb. Optional. Video thumbnail.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileSize. Optional. File size bytes.
	FileSize int `json:"file_size,omitempty"`
}
