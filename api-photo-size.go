package tgwrap

//
// Telegram API object
//
// PhotoSize represents one size of a photo or a file / sticker thumbnail.
//
type PhotoSize struct {

	//
	// Unique identifier for this file
	//
	FileID string `json:"file_id"`

	//
	// Photo width
	//
	Width uint `json:"width"`

	//
	// Photo height
	//
	Height uint `json:"height"`

	//
	// Optional. File size
	//
	FileSize uint `json:"file_size,omitempty"`
}
