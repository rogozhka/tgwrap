package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputFileLocal(t *testing.T) {

	path := "api-input-file.go"

	f := NewInputFileLocal(path)

	assert.Equal(t, path, f.String(), "Name is set")
}
