package tgwrap

// ChatPhoto is a Telegram API object.
type ChatPhoto struct {

	// SmallFileID - unique file identifier of small (160x160) chat photo.
	// This file_id can be used only for photo download.
	SmallFileID string `json:"small_file_id"`

	// SmallFileUniqueID - unique file identifier of small (160x160) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id"`

	// BigFileID - unique file identifier of big (640x640) chat photo.
	// This file_id can be used only for photo download.
	BigFileID string `json:"big_file_id"`

	// BigFileUniqueID - unique file identifier of big (640x640) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id"`
}
