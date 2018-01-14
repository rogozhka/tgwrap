package thestruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTags_index(t *testing.T) {

	tag := `json:"name,string" xml:"title" gorm:"primary_key" bigquery:",nullable"`

	tags, err := ParseLiteral(tag)

	assert.Nil(t, err, "ParseLiteral")
	assert.NotNil(t, tags, "notnil tags")

	assert.Equal(t, 4, len(tags.Sections()), "len arrTags ")

	json := tags.Tag("json")
	assert.NotNil(t, json, "notnil tag")

}
