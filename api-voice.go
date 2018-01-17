package tgwrap

//
// Voice represents a voice note.
//
type Voice struct {

	//
	// Unique identifier for this file
	//
	ID string `json:"file_id"`

	//
	// Duration of the audio in seconds as defined by sender
	//
	Duration uint `json:"duration"`

	//
	// Optional. MIME type of the file as defined by sender
	//
	MIMEType string `json:"mime_type,omitempty"`

	//
	// Optional. File size
	//
	FileSize uint `json:"file_size,omitempty"`
}
