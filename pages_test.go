package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateFilename(t *testing.T) {
	assert.True(t, ValidateFilename("test.md"))
	assert.True(t, ValidateFilename("test"))
	assert.True(t, ValidateFilename("test.jpg"))
	assert.True(t, ValidateFilename("test.A"))
	assert.True(t, ValidateFilename("test.aa.AA.aa"))
	assert.True(t, ValidateFilename("test_aa"))
	assert.True(t, ValidateFilename("test-TEST"))

	assert.False(t, ValidateFilename("t est.md"))
	assert.False(t, ValidateFilename("test .md"))
	assert.False(t, ValidateFilename("t?est.md"))
	assert.False(t, ValidateFilename("test!"))
	assert.False(t, ValidateFilename("test%^"))
}
