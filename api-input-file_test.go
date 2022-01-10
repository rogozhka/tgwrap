package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputFileLocal(t *testing.T) {
	is := assert.New(t)
	path := "api-input-file.go"
	f := NewInputFileLocal(path)
	is.Equal(path, f.String(), "String()")
}
func TestNewInputFileFromID(t *testing.T) {
	is := assert.New(t)
	id := "test-file-id"
	f := NewInputFileFromID(id)
	is.Equal(id, f.String(), "String()")
}
func TestNewInputFileFromURL(t *testing.T) {
	is := assert.New(t)
	url := "http://example.com/test-file-path"
	f := NewInputFileFromID(url)
	is.Equal(url, f.String(), "String()")
}
