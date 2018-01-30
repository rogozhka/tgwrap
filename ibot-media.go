package tgwrap

//
// IBotMedia is a group of methods for sending media
// used as a part of common IBot API interface
//
type IBotMedia interface {
	//
	// SendPhoto is used to send photos.
	//
	// chatID: (Integer or String) Unique identifier
	// for the target chat or username of the target channel (in the format @channelusername)
	//
	// photo: (*InputFile or string) Photo to send. Pass a file_id as string to send
	// a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a string
	// for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data
	// using &NewInputFileLocal("<file path>")
	//
	// opt: (can be nil) optional params
	//
	SendPhoto(chatID interface{}, photo interface{}, opt *SendPhotoOpt) (*Message, error)

	//
	// SendVideo is used to send video files, Telegram clients support mp4 videos
	// (other formats may be sent as Document). On success, the sent Message is returned.
	// Bots can currently send video files of up to 50 MB in size,
	// this limit may be changed in the future.
	//
	// chatID: (uint64 or string) Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// video: (*InputFile or string) Video to send. Pass a file_id as string to send
	// a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a string
	// for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data
	// using &NewInputFileLocal("<file path>")
	//
	// opt: (can be nil) optional params
	//
	SendVideo(chatID interface{}, video interface{}, opt *SendVideoOpt) (*Message, error)

	// SendVideoNote is used to send video notes.
	// As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long.
	// Use this method to send video messages. On success, the sent Message is returned.
	//
	// chatID: (uint64 or string) Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// video: (*InputFile or string) Video note to send.
	// Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or
	// upload a new video using multipart/form-data.
	// Sending video notes by a URL is currently unsupported
	//
	// opt: (can be nil) optional params
	//
	SendVideoNote(chatID interface{}, video interface{}, opt *SendVideoNoteOpt) (*Message, error)

	//
	// SendAudio is used to send audio files, if you want Telegram clients to display them in the music player.
	// Your audio must be in the .mp3 format. On success, the sent Message is returned. Bots can currently
	// send audio files of up to 50 MB in size, this limit may be changed in the future.
	//
	// For sending voice messages, use the SendVoice method instead.
	//
	// chatID: (uint64 or string) Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// audio: (*InputFile or string) Audio to send. Pass a file_id as string to send
	// an audio that exists on the Telegram servers (recommended), pass an HTTP URL as a string
	// for Telegram to get an audio from the Internet, or upload a new file using multipart/form-data.
	// using &NewInputFileLocal("<file path>")
	//
	// opt: (can be nil) optional params
	//
	SendAudio(chatID interface{}, audio interface{}, opt *SendAudioOpt) (*Message, error)

	//
	// SendVoice to send audio files, if you want Telegram clients to display the file
	// as a playable voice message. For this to work, your audio must be in an .ogg file
	// encoded with OPUS (other formats may be sent as Audio or Document).
	// On success, the sent Message is returned. Bots can currently send voice messages
	// of up to 50 MB in size, this limit may be changed in the future.
	//
	// chatID: (uint64 or string) Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// voice: (*InputFile or string) Audio to send. Pass a file_id as String to send
	// an audio that exists on the Telegram servers (recommended), pass an HTTP URL as a String
	// for Telegram to get an audio from the Internet, or upload a new file using multipart/form-data.
	// using &NewInputFileLocal("<file path>")
	//
	// opt: (can be nil) optional params
	//
	SendVoice(chatID interface{}, voice interface{}, opt *SendVoiceOpt) (*Message, error)

	//
	// SendDocument to send general files. On success, the sent Message is returned.
	// Bots can currently send files of any type of up to 50 MB in size, t
	// his limit may be changed in the future.
	//
	// chatID: (uint64 or string) Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// document: (*InputFile or string) File to send. Pass a file_id as String to send
	// an audio that exists on the Telegram servers (recommended), pass an HTTP URL as a String
	// for Telegram to get an audio from the Internet, or upload a new file using multipart/form-data.
	// using &NewInputFileLocal("<file path>")
	//
	// opt: (can be nil) optional params
	//
	SendDocument(chatID interface{}, document interface{}, opt *SendDocumentOpt) (*Message, error)

	// SendLocation is used to send point on the map.
	// On success, the sent Message is returned.
	//
	// chatID: (uint64 or string) Unique identifier for the target chat
	// or username of the target channel (in the format @channelusername)
	//
	// latitude: (float64) Latitude of the location.
	//
	// longitude: (float64) Longitude of the location.
	//
	// opt: (can be nil) optional params
	//
	SendLocation(chatID interface{}, latitude float64, longitude float64, opt *SendLocationOpt) (*Message, error)
}
