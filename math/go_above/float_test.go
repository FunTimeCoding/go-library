package go_above

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assertFloat(t, 52, 51, 50, false)
	assertFloat(t, 51, 50, 50, false)
	assertFloat(t, 51, 49, 50, false)
	assertFloat(t, 49, 51, 50, true)
}

func assertFloat(
	t *testing.T,
	past float64,
	now float64,
	threshold float64,
	expected bool,
) {
	t.Helper()
	assert.Boolean(t, expected, Float(past, now, threshold))
}
