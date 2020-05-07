package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBot_GetFile(t *testing.T) {

	bot := createTestBotFromEnv()

	fileID := testEnvStringValue("TGWRAP_TEST_FILE_ID")

	file, err := bot.GetFile(fileID)
	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.True(t, len(file.FilePath) > 0)
	assert.True(t, file.FileSize > 0)
}
