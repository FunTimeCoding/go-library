package above_below

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	assertInteger(t, 1, 0, true, false)
	assertInteger(t, -1, 0, false, true)
	assertInteger(t, 0, 0, false, false)
}

func assertInteger(
	t *testing.T,
	f int,
	magnitude int,
	expectedAbove bool,
	expectedBelow bool,
) {
	t.Helper()
	var above bool
	var below bool
	Integer(
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
