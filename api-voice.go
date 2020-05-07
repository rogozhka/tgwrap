package tgwrap

// Voice represents a voice note.
type Voice struct {

	// ID is unique identifier for this file.
	ID string `json:"file_id"`

	// Duration of the audio in seconds as defined by sender.
	Duration int `json:"duration"`

	// MIMEType. Optional. MIME type of the file as defined by sender.
	MIMEType string `json:"mime_type,omitempty"`

	// FileSize. Optional. File size.
	FileSize int `json:"file_size,omitempty"`
}
