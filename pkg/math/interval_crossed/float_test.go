package interval_crossed

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assert2.False(t, Float(10, 0, 1))
	assert2.False(t, Float(10, 1, 9))
	assert2.True(t, Float(10, 9, 10))
	assert2.True(t, Float(10, 9, 11))
	assert2.False(t, Float(10, 10, 11))
}
