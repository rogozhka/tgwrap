package tgwrap

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {

	// FileID is identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID - unique identifier for this file, which is supposed
	// to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by sender.
	Duration int `json:"duration"`

	// Performer is optional performer name of the audio as defined by sender or by audio tags.
	Performer string `json:"performer,omitempty"`

	// Title is optional title of the audio as defined by sender or by audio tags.
	Title string `json:"title,omitempty"`

	// MIMEType is optional type of the file as defined by sender.
	MIMEType string `json:"mime_type,omitempty"`

	// FileSize is file size in bytes if known.
	FileSize int `json:"file_size,omitempty"`

	// Thumb is optional thumbnail of the album cover to which the music file belongs.
	Thumb *PhotoSize `json:"thumb,omitempty"`
}
