package tgwrap

import (
	"crypto/md5"
	"testing"
	"time"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test_SetChatTitle(t *testing.T) {
	bot := createTestBotFromEnv()

	chatID := requireEnv(envTestChatID)

	newTitle := fmt.Sprintf("Title: %v", md5.Sum([]byte(time.Now().String())))

	_, err := bot.SetChatTitle(chatID, newTitle)

	assert.Nil(t, err, "SetChatTitle err")
}
