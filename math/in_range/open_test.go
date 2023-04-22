package in_range

import (
	"github.com/funtimecoding/go-library/assert"
	"github.com/funtimecoding/go-library/math/ranges"
	"testing"
)

func TestOpen(t *testing.T) {
	zeroToOne := ranges.Range{L: 0, R: 1}
	assertOpen(t, 0, zeroToOne, false)
	assertOpen(t, 0.01, zeroToOne, true)
	assertOpen(t, 0.99, zeroToOne, true)
	assertOpen(t, 1, zeroToOne, false)
	assertOpen(t, 1.01, zeroToOne, false)
}

func assertOpen(
	t *testing.T,
	value float64,
	r ranges.Range,
	expected bool,
) {
	t.Helper()
	assert.Boolean(t, Open(value, r), expected)
}
