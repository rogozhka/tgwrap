package thestruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_explodeStructTag(t *testing.T) {

	tag := `json:"name,string" xml:"title" gorm:"primary_key" bigquery:",nullable"`

	arrTags, err := explodeStructTag(tag)

	assert.Nil(t, err, "parse")
	assert.Equal(t, 4, len(arrTags), "len arrTags ")

	assert.Equal(t, "json", arrTags[0].Section, "Section")
	assert.Equal(t, "name", arrTags[0].Value, "Value")
	assert.True(t, arrTags[0].IsOpt("string"), "Opt")
	assert.Equal(t, 1, len(arrTags[0].Options), "Opt")

	assert.Equal(t, "xml", arrTags[1].Section, "Section")
	assert.Equal(t, "title", arrTags[1].Value, "Value")
	assert.False(t, arrTags[1].IsOpt("string"), "Opt")
	assert.Equal(t, 0, len(arrTags[1].Options), "Opt")

	assert.Equal(t, "bigquery", arrTags[3].Section, "Section")
	assert.Equal(t, "", arrTags[3].Value, "Value")
	assert.False(t, arrTags[3].IsOpt("string"), "Opt")
	assert.True(t, arrTags[3].IsOpt("nullable"), "Opt")
	assert.Equal(t, 1, len(arrTags[3].Options), "Opt")
}
