package channel

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestChannel(t *testing.T) {
	assert.True(t, New() != nil)
}
