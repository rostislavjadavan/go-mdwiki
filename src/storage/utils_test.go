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

func TestParsePageFilename(t *testing.T) {
	f, ts := parsePageFilename("test1.md")
	assert.Equal(t, "test1.md", f)
	assert.Equal(t, int64(0), ts)

	f, ts = parsePageFilename("test5.md__105")
	assert.Equal(t, "test5.md", f)
	assert.Equal(t, int64(105), ts)

	f, ts = parsePageFilename("test__166.md__99")
	assert.Equal(t, "test__166.md", f)
	assert.Equal(t, int64(99), ts)

	f, ts = parsePageFilename("test.md__66.md__11")
	assert.Equal(t, "test.md__66.md", f)
	assert.Equal(t, int64(11), ts)

	f, ts = parsePageFilename("test.md__aa")
	assert.Equal(t, "test.md", f)
	assert.Equal(t, int64(0), ts)
}
