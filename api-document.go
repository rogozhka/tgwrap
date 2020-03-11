package tgwrap

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {

	// FileID is unique file identifier.
	FileID string `json:"file_id"`

	// Thumb. Optional. Document thumbnail as defined by sender.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileName. Optional. Original filename as defined by sender.
	FileName string `json:"file_name,omitempty"`

	// MIMEType. Optional. MIME type of the file as defined by sender.
	MIMEType string `json:"mime_type,omitempty"`

	// FileSize. Optional. File size bytes.
	FileSize uint `json:"file_size,omitempty"`

	// FilePath. Optional. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath string `json:"file_path,omitempty"`
}
