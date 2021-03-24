package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHighlight(t *testing.T) {
	assert.Equal(t, "aa <b class=\"highlight\">BB</b> dd", highlight("aa BB dd", "BB", 0, "bb"))
}
