package fall_below

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	// Short by 1
	assertFloat(t, 51, 50, 50, false)

	// Reached exactly
	assertFloat(t, 50, 49, 50, true)

	// Exceed by 1
	assertFloat(t, 51, 49, 50, true)

	// Go above
	assertFloat(t, 49, 51, 50, false)
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
