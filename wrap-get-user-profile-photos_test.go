package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBot_GetUserProfilePhotos(t *testing.T) {

	bot := createTestBotFromEnv()

	userID := int64(testEnvIntValue(envTestUserID))

	res, err := bot.GetUserProfilePhotos(userID, nil)
	assert.Nil(t, err, userID)
	assert.NotNil(t, res)
	assert.True(t, res.TotalCount > 0)
	assert.Equal(t, res.TotalCount, len(res.Photos))
}
