package tunnel

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTunnel(t *testing.T) {
	assert.True(t, New() != nil)
}
