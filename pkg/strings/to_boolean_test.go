package strings

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToBoolean(t *testing.T) {
	assert2.True(t, ToBoolean("1"))
	assert2.False(t, ToBoolean("0"))
}
