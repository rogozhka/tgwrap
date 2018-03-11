package tgwrap

import (
	"crypto/md5"
	"testing"
	"time"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test_SetChatTitle(t *testing.T) {
	token := getTokenEnv()
	bot := createBot(token)
	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")

	newTitle := fmt.Sprintf("Title: %v", md5.Sum([]byte(time.Now().String())))

	_, err4 := bot.SetChatTitle(chatID, newTitle)

	assert.Nil(t, err4, "SetChatTitle err")
}
