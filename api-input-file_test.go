package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputFile__ReadFromLocal(t *testing.T) {

	path := "api-input-file.go"

	f := NewInputFileLocal(path)

	assert.Equal(t, path, f.Name(), "Name is set")
}
