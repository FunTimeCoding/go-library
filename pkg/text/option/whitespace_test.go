package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestWhitespace(t *testing.T) {
	assert.True(t, New() != nil)
}
