package channel

import (
	"testing"

	"github.com/funtimecoding/go-library/pkg/assert"
)

func TestChannel(t *testing.T) {
	assert.True(t, New() != nil)
}
