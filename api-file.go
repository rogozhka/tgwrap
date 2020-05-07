package tgwrap

// File represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour. When the link expires,
// a new one can be requested by calling getFile. Maximum file size to download is 20 MB.
type File struct {

	// FileID is identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID - unique identifier for this file, which is supposed
	// to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// FileSize in bytes, if known.
	FileSize int `json:"file_size,omitempty"`

	// FilePath - use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath string `json:"file_path"`
}
