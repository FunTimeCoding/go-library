package normalize_change

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assertFloat(t, 0, 1, 0, 100, 1)
	assertFloat(t, 1, -2, 0, 100, -1)
	assertFloat(t, 1, -3, 0, 100, -1)
	assertFloat(t, 100, -100, 0, 100, -100)
	assertFloat(t, 100, -150, 0, 100, -100)
}

func TestFloatNoMaximum(t *testing.T) {
	assertFloat(t, 0, 1, 0, 0, 1)
	assertFloat(t, 1, -2, 0, 0, -1)
	assertFloat(t, 1, -3, 0, 0, -1)
	assertFloat(t, 100, -100, 0, 0, -100)
	assertFloat(t, 100, -150, 0, 0, -100)
}

func assertFloat(
	t *testing.T,
	now float64,
	change float64,
	minimum float64,
	maximum float64,
	expected float64,
) {
	t.Helper()
	assert.Float(t, expected, Float(now, change, minimum, maximum))
}
