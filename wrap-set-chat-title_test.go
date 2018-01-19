package tgwrap

import (
	"crypto/md5"
	"os"
	"testing"
	"time"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test_SetChatTitle(t *testing.T) {
	token, err := getTokenEnv()
	assert.Nil(t, err, "Token")
	if len(token) < 1 {
		return
	}
	bot := &bot{
		token: token,
	}

	envName := "TGWRAP_TEST_CHAT_ID"
	chatID := os.Getenv(envName)
	if len(chatID) < 1 {
		return
	}

	newTitle := fmt.Sprintf("Title: %v", md5.Sum([]byte(time.Now().String())))

	_, err4 := bot.SetChatTitle(chatID, newTitle)

	assert.Nil(t, err4, "SetChatTitle err")
}
