package round_tripper

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRoundTripper(t *testing.T) {
	assert.True(t, New("", "") != nil)
}
