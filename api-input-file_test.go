package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputFileLocal(t *testing.T) {

	path := "api-input-file.go"

	f := NewInputFileLocal(path)

	assert.Equal(t, path, f.String(), "String()")
}

func TestNewInputFileFromID(t *testing.T) {

	id := "test-file-id"
	f := NewInputFileFromID(id)

	assert.Equal(t, id, f.String(), "String()")
}

func TestNewInputFileFromURL(t *testing.T) {

	url := "http://example.com/test-file-path"
	f := NewInputFileFromID(url)

	assert.Equal(t, url, f.String(), "String()")
}
