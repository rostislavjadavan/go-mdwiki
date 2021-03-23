package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateFilename(t *testing.T) {
	assert.Nil(t, ValidateFilename("test.md"))
	assert.Nil(t, ValidateFilename("test"))
	assert.Nil(t, ValidateFilename("test.jpg"))
	assert.Nil(t, ValidateFilename("test.A"))
	assert.Nil(t, ValidateFilename("test.aa.AA.aa"))
	assert.Nil(t, ValidateFilename("test_aa"))
	assert.Nil(t, ValidateFilename("test-TEST"))

	assert.NotNil(t, ValidateFilename("t est.md"))
	assert.NotNil(t, ValidateFilename("test .md"))
	assert.NotNil(t, ValidateFilename("t?est.md"))
	assert.NotNil(t, ValidateFilename("test!"))
	assert.NotNil(t, ValidateFilename("test%^"))
}

func TestFixPageExtension(t *testing.T) {
	assert.Equal(t, "", FixPageExtension(""))
	assert.Equal(t, "test.md", FixPageExtension("test"))
	assert.Equal(t, "test.md", FixPageExtension("test.md"))
}
