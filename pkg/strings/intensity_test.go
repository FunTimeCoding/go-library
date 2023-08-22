package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestIntensity(t *testing.T) {
	elements := []string{
		Alfa,
		Bravo,
		Charlie,
	}

	assert.String(t, "Alfa", Intensity(elements, 0))
	assert.String(t, "Alfa", Intensity(elements, 0.33))

	assert.String(t, "Bravo", Intensity(elements, 0.34))
	assert.String(t, "Bravo", Intensity(elements, 0.66))

	assert.String(t, "Charlie", Intensity(elements, 0.67))
	assert.String(t, "Charlie", Intensity(elements, 1))
}
