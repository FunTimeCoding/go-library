package in_range

import (
	"github.com/funtimecoding/go-library/assert"
	"github.com/funtimecoding/go-library/math/ranges"
	"testing"
)

func TestRightOpen(t *testing.T) {
	zeroToOne := ranges.Range{L: 0, R: 1}
	assertRightOpen(t, 0, zeroToOne, true)
	assertRightOpen(t, 0.01, zeroToOne, true)
	assertRightOpen(t, 0.99, zeroToOne, true)
	assertRightOpen(t, 1, zeroToOne, false)
	assertRightOpen(t, 1.01, zeroToOne, false)
}

func assertRightOpen(
	t *testing.T,
	value float64,
	r ranges.Range,
	expected bool,
) {
	t.Helper()
	assert.Boolean(t, RightOpen(value, r), expected)
}
