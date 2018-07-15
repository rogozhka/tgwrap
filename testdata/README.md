# Testing
```bash

#
# Assume you have a project
# and imported tgwrap by `go get -u github.com/rogozhka/tgwrap`
#
go test ./src/github.com/rogozhka/tgwrap/

```



## Required ENV variables



```bash export TGWRAP_TOKEN="<your test bot token>"

#
# Register test bot with @BotFather, obtain token
# set proper chat options, picture, description.
#
# Note: never use real bots in prod for automated tests
#
export TGWRAP_TOKEN="your test bot's access token"

#
# Test bot should have admin privileges in group chat
# and chat option 'All member are admins' = FALSE
#
export TGWRAP_TEST_CHAT_ID="chatID for testing"

#
# User ID for testing individual messages
#
export TGWRAP_TEST_PERSONAL_CHAT_ID="chatID for testing"


#
# Where to get test video
#
# Note: Telegram supports only mp4 format
# if you want video playback teaser and thumbnail
#
export TGWRAP_DOWNLOAD_URL_VIDEO="http://<mp4 video location>"

#
# File name to use for upload video
#
export TGWRAP_TEST_VIDEO_PATH="video.mp4"



```

Alternatively you can setup these variables in Travis control panel.









