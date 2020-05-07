package tgwrap

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {

	// FileID - unique identifier for this file.
	FileID string `json:"file_id"`

	// Width is Photo width.
	Width int `json:"width"`

	// Height is Photo height.
	Height int `json:"height"`

	// Optional. File size.
	FileSize int `json:"file_size,omitempty"`
}
