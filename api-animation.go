package tgwrap

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
// You can provide an animation for your game so that
// it looks stylish in chats (check out Lumberjack for an example).
// This object represents an animation file to be displayed
// in the message containing a game.
type Animation struct {

	// FileID is identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// Thumb is optional animation thumbnail as defined by sender.
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// FileName is optional original animation filename as defined by sender.
	FileName string `json:"file_name,omitempty"`

	// MIMEType is optional type of the file as defined by sender.
	MIMEType string `json:"mime_type,omitempty"`

	// FileSize is file size in bytes if known.
	FileSize int `json:"file_size,omitempty"`
}
