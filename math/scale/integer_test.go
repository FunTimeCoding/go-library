package scale

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	assertInteger(t, 0, 100, 0, 0)
	assertInteger(t, 0, 100, 0.5, 50)
	assertInteger(t, 0, 100, 1, 100)
	assertInteger(t, 40, 20, 0.5, 30)
}

func assertInteger(
	t *testing.T,
	from int,
	to int,
	factor float64,
	expected int,
) {
	t.Helper()
	assert.Integer(t, expected, Integer(from, to, factor))
}
