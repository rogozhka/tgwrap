package tgwrap

type IBotMedia interface {

	//
	// Use this method to send photos.
	//
	// @param chatID (Integer or String) Unique identifier
	// for the target chat or username of the target channel (in the format @channelusername)
	//
	// @param photo (InputFile or String) Photo to send. Pass a file_id as String to send
	// a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String
	// for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data.
	//
	SendPhoto(chatID interface{}, photo interface{}, opt *SendPhotoOpt) (*Message, error)
}
