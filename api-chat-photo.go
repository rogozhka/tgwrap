package tgwrap

//
// Telegram API object
//
type ChatPhoto struct {

	//
	// Unique file identifier of small (160x160) chat photo.
	// This file_id can be used only for photo download.
	//
	SmallFileID string `json:"small_file_id"`

	//
	// Unique file identifier of big (640x640) chat photo.
	// This file_id can be used only for photo download.
	//
	BigFileID string `json:"big_file_id"`
}
