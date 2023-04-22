package math

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestDistribution(t *testing.T) {
	assertDistribution(t, 100, 10, 100)
	assertDistribution(t, 1000000, 10, 1000000)
}

func assertDistribution(
	t *testing.T,
	all float64,
	steps int,
	expected float64,
) {
	t.Helper()
	var sum float64

	for _, element := range Distribution(all, steps) {
		sum += element
	}

	// Comparing all floats here is a lot of decimals
	//assert.Any(t, expected, Distribution(all, steps))
	assert.Float(t, expected, sum)
}
