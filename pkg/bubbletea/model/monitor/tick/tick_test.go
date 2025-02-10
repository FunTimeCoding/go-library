package tick

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTick(t *testing.T) {
	assert.True(t, Command() != nil)
}
