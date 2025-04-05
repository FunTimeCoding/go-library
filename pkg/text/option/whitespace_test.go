package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestWhitespace(t *testing.T) {
	c := Compact()
	assert.False(t, c.NewlineAtEnd)
	assert.Integer(t, 0, c.AllowedBlankLines)
}
