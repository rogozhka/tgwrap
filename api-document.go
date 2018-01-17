package tgwrap

//
// Document represents a general file
// (as opposed to photos, voice messages and audio files).
//
type Document struct {

	//
	// Unique file identifier
	//
	FileID string `json:"file_id"`

	//
	// Optional. Document thumbnail as defined by sender
	//
	Thumb *PhotoSize `json:"thumb,omitempty"`

	//
	// Optional. Original filename as defined by sender
	//
	FileName string `json:"file_name,omitempty"`

	//
	// Optional. MIME type of the file as defined by sender
	//
	MIMEType string `json:"mime_type,omitempty"`

	//
	// Optional. File size
	//
	FileSize uint `json:"file_size,omitempty"`
}
