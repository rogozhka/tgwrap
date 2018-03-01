package tgwrap

// This object represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile.
type File struct {

	//
	// Unique identifier for this file
	//
	FileID string `json:"file_id"`

	//
	// Optional. File size, if known
	//
	FileSize int64 `json:"file_size"`

	//
	// Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	//
	FilePath string `json:"file_path"`
}
