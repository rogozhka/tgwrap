package tgwrap

//
// Audio represents an audio file to be treated as music by the Telegram clients.
//
type Audio struct {

	//
	// Unique identifier for this file
	//
	FileID string `json:"file_id"`

	//
	// Duration of the audio in seconds as defined by sender
	//
	Duration uint `json:"duration"`

	//
	// Optional. Performer of the audio as defined by sender or by audio tags
	//
	Performer string `json:"performer,omitempty"`

	//
	// Optional. Title of the audio as defined by sender or by audio tags
	//
	Title string `json:"title,omitempty"`

	//
	// Optional. MIME type of the file as defined by sender
	//
	MIMEType string `json:"mime_type,omitempty"`

	//
	// Optional. File size
	//
	FileSize uint `json:"file_size,omitempty"`

	//
	// Optional. Thumbnail of the album cover to which the music file belongs
	//
	Thumb *PhotoSize `json:"thumb,omitempty"`
}
