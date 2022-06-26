package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBot_GetFile(t *testing.T) {
	is := assert.New(t)
	bot := createTestBotFromEnv()

	fileID := testEnvStringValue(envTestFileID)

	file, err := bot.GetFile(fileID)
	is.Nil(err)
	is.NotNil(file)
	is.True(len(file.FilePath) > 0)
	is.True(file.FileSize > 0)
}
