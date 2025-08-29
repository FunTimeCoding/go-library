package ptr

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTo(t *testing.T) {
	assert.NotNil(t, To(0))
}
