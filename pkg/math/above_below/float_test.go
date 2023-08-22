package above_below

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assertFloat(t, 1, 0, true, false)
	assertFloat(t, -1, 0, false, true)
	assertFloat(t, 0, 0, false, false)

	assertFloat(t, 1, 1, false, false)
	assertFloat(t, -1, 1, false, false)
	assertFloat(t, 0, 1, false, false)

	assertFloat(t, 2, 1, true, false)
	assertFloat(t, -2, 1, false, true)
}

func assertFloat(
	t *testing.T,
	f float64,
	magnitude float64,
	expectedAbove bool,
	expectedBelow bool,
) {
	t.Helper()
	var above bool
	var below bool
	Float(
		f,
		magnitude,
		func() {
			above = true
		},
		func() {
			below = true
		},
	)
	assert.Boolean(t, expectedAbove, above)
	assert.Boolean(t, expectedBelow, below)
}
