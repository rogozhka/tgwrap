package tgwrap

type IBotMedia interface {

	//
	// SendPhoto to send photos.
	//
	// @param chatID (Integer or String) Unique identifier
	// for the target chat or username of the target channel (in the format @channelusername)
	//
	// @param photo (InputFile or String) Photo to send. Pass a file_id as String to send
	// a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String
	// for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data.
	//
	SendPhoto(chatID interface{}, photo interface{}, opt *SendPhotoOpt) (*Message, error)

	//
	// SendVideo to send video files, Telegram clients support mp4 videos
	// (other formats may be sent as Document). On success, the sent Message is returned.
	// Bots can currently send video files of up to 50 MB in size,
	// this limit may be changed in the future.
	//
	// @param chatID (Integer or String) Unique identifier
	// for the target chat or username of the target channel (in the format @channelusername)
	//
	// @param video (InputFile or String) Video to send. Pass a file_id as String to send
	// a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String
	// for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data.
	//
	SendVideo(chatID interface{}, video interface{}, opt *SendVideoOpt) (*Message, error)

	//
	// SendAudio to send audio files, if you want Telegram clients to display them in the music player.
	// Your audio must be in the .mp3 format. On success, the sent Message is returned. Bots can currently
	// send audio files of up to 50 MB in size, this limit may be changed in the future.
	//
	// For sending voice messages, use the sendVoice method instead.
	//
	// @param audio (InputFile or String) Audio to send. Pass a file_id as String to send
	// a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String
	// for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data.
	//
	SendAudio(chatID interface{}, audio interface{}, opt *SendAudioOpt) (*Message, error)

	//
	// SendVoice to send audio files, if you want Telegram clients to display the file
	// as a playable voice message. For this to work, your audio must be in an .ogg file
	// encoded with OPUS (other formats may be sent as Audio or Document).
	// On success, the sent Message is returned. Bots can currently send voice messages
	// of up to 50 MB in size, this limit may be changed in the future.
	//
	SendVoice(chatID interface{}, voice interface{}, opt *SendVoiceOpt) (*Message, error)
}
