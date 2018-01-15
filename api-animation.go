package tgwrap

//
// Telegram API object
//
// You can provide an animation for your game so that
// it looks stylish in chats (check out Lumberjack for an example).
// This object represents an animation file to be displayed
// in the message containing a game.
//
type Animation struct {

	//
	// Unique file identifier
	//
	FileID string `json:"file_id"`

	//
	// Optional. Animation thumbnail as defined by sender
	//
	Thumb *PhotoSize `json:"thumb,omitempty"`

	//
	// Optional. Original animation filename as defined by sender
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
