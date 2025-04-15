package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestIntensity(t *testing.T) {
	v := []string{Alfa, Bravo, Charlie}

	assert.String(t, "Alfa", Intensity(v, 0))
	assert.String(t, "Alfa", Intensity(v, 0.33))

	assert.String(t, "Bravo", Intensity(v, 0.34))
	assert.String(t, "Bravo", Intensity(v, 0.66))

	assert.String(t, "Charlie", Intensity(v, 0.67))
	assert.String(t, "Charlie", Intensity(v, 1))
}
