package reach_reduce

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	// Short by 1
	assertInteger(t, 52, 51, 50, false)

	// Reached exactly
	assertInteger(t, 51, 50, 50, true)

	// Exceed by 1
	assertInteger(t, 51, 49, 50, true)

	// Reach increase
	assertInteger(t, 49, 51, 50, false)
}

func assertInteger(
	t *testing.T,
	past int,
	now int,
	threshold int,
	expected bool,
) {
	t.Helper()
	assert.Boolean(t, expected, Integer(past, now, threshold))
}
