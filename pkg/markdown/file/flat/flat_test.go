package flat

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFlat(t *testing.T) {
	assert.True(t, New() != nil)
}
