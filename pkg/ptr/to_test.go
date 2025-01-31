package ptr

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTo(t *testing.T) {
	assert.True(t, To(0) != nil)
}
