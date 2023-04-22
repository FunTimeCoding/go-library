package normalize

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	// Meet minimum
	assertFloat(t, 0, 0, 100, 0)
	// Below minimum
	assertFloat(t, -1, 0, 100, 0)
	// Meet maximum
	assertFloat(t, 100, 0, 100, 100)
	// Above maximum
	assertFloat(t, 101, 0, 100, 100)
	// No maximum
	assertFloat(t, 101, 0, 0, 101)
}

func assertFloat(
	t *testing.T,
	f float64,
	minimum float64,
	maximum float64,
	expected float64,
) {
	t.Helper()
	Float(&f, minimum, maximum)
	assert.Round(t, expected, f, 0)
}
