package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToBoolean(t *testing.T) {
	assert.True(t, ToBoolean("1"))
	assert.False(t, ToBoolean("0"))
}
