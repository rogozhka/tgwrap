package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBot_GetUserProfilePhotos(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	userID := int64(testEnvIntValue(envTestUserID))

	res, err := bot.GetUserProfilePhotos(userID, nil)
	is.Nil(err, userID)
	is.NotNil(res)
	is.True(res.TotalCount > 0)
	is.Equal(res.TotalCount, len(res.Photos))
}
