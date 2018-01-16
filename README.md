[![GoDoc](https://godoc.org/github.com/rogozhka/tgwrap?status.svg)](https://godoc.org/github.com/rogozhka/tgwrap)
[![Travis](https://travis-ci.org/rogozhka/tgwrap.svg?branch=master)](https://travis-ci.org/rogozhka/tgwrap)

**TgWrap** is a Go library with [Telegram Bot API](https://core.telegram.org/bots/api) bindings. General approach is to keep objects as close as possible to the API structures with respect to optional params. Results of the methods are the same as noted in API documentation. Library does not pretend to be framework with tons of handlers and application strategy in mind. But there are some private helpers w/ reflection and struct tags under the hood. The aim is to cover step by step all the methods including Payment section.



## Example

```go
package main

import (
	"log"
	"fmt"
	"github.com/rogozhka/tgwrap"
)

func main() {

	bot := tgwrap.NewBot("<token>")

  	// returns *User
	me, err := bot.GetMe()
	if err != nil {
		panic(err)
	}

	log.Printf("BotID:%v", me.ID)
	log.Printf("BotName:%v", me.FirstName)

	updates, err := bot.GetUpdates(&tgwrap.GetUpdatesOpt{
		Timeout: 60,
	})
	if err != nil {
		panic(err)
	}

	for _, u := range updates {

		log.Printf("UpdateID:%v", u.ID)
		log.Printf("MessageID:%v", u.Message.ID)
		log.Printf("ChatID:%v", u.Message.Chat.ID)
		log.Printf("Text:%v", u.Message.Text)
		log.Printf("UserID:%v", u.Message.From.ID)
		log.Printf("FirstName:%v", u.Message.From.FirstName)

		// reply to message with "Re: <original message>"
		bot.SendMessage(u.Message.Chat.ID, fmt.Sprintf("Re: %v", u.Message.Text),
			&tgwrap.SendMessageOpt{
				DisableNotification: true,
			})

		// send photo from URL
		bot.SendPhoto(u.Message.Chat.ID,
             tgwrap.NewInputFileFromURL("http://imageurl.com/test.jpg"),
			&tgwrap.SendPhotoOpt{
				DisableNotification: true,
				Caption:             "Test gray photo",
			})

		bot.SendChatAction(u.Message.Chat.ID, tgwrap.ChatActionUploadPhoto)

		// send photo from fs
		bot.SendPhoto(u.Message.Chat.ID,
			tgwrap.NewInputFileLocal("/tmp/local-file-name.jpg"),
			&tgwrap.SendPhotoOpt{
				DisableNotification: true,
				Caption:             "Photo from webcam ;)",
			})

		bot.SendChatAction(u.Message.Chat.ID, tgwrap.ChatActionRecordAudio)

		// send audio from fs
		bot.SendAudio(u.Message.Chat.ID,
			tgwrap.NewInputFileLocal("/tmp/song.mp3"),
			&tgwrap.SendAudioOpt{
				DisableNotification: true,
				Title:               "Song name",
				Performer:           "The Family Band",
				Caption:             "Lastest record",
			})
	}
}
```



## Currently supported methods

[IBotUpdates](https://godoc.org/github.com/rogozhka/tgwrap#IBotUpdates) to get updates and setup hooks:

- [x] [getUpdates](https://core.telegram.org/bots/api#getupdates)

- [ ] setWebhook

- [ ] deleteWebhook

- [ ] getWebhookInfo





[IBotMessages](https://godoc.org/github.com/rogozhka/tgwrap#IBotMessages) to send messages and chat actions:

- [x] [sendMessage](https://core.telegram.org/bots/api#sendmessage)
- [x] [forwardMessage](https://core.telegram.org/bots/api#forwardmessage)




[IBotChat](https://godoc.org/github.com/rogozhka/tgwrap#IBotChat) to interact with chat state and settings:

- [x] [sendChatAction](https://core.telegram.org/bots/api#sendchataction)

- [ ] ...




[IBotMedia](https://godoc.org/github.com/rogozhka/tgwrap#IBotMedia) for sending media objects, encoding local files, by url, file_id:

- [x] [sendPhoto](https://core.telegram.org/bots/api#sendphoto)

- [x] [sendAudio](https://core.telegram.org/bots/api#sendaudio)

- [x] [sendVideo](https://core.telegram.org/bots/api#sendvideo)

- [x] [sendVoice](https://core.telegram.org/bots/api#sendvoice)

- [ ] sendDocument

- [ ] sendVideoNote

- [ ] sendMediaGroup

- [ ] ...



## Contribution

Welcome feedback Issues and PR.



## Licence

Released under the [MIT License](https://github.com/rogozhka/tgwrap/blob/master/LICENSE).