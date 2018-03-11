package tgwrap

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
)

func Test_UnbanChatMember(t *testing.T) {

	token := getTokenEnv()
	bot := createBot(token)
	chatID := requireEnv("TGWRAP_TEST_CHAT_ID")
	strUserID := requireEnv("TGWRAP_TEST_USER_ID")

	return // till decision how to test supergroups

	userID, err := strconv.Atoi(strUserID)
	assert.Nil(t, err, "Atoi err")

	_, err4 := bot.UnbanChatMember(chatID, uint64(userID))
	assert.Nil(t, err4, "UnbanChatMember err")
}
