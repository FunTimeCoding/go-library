package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestToLower(t *testing.T) {
	assert.Strings(
		t,
		[]string{"alfa", "bravo"},
		ToLower([]string{Alfa, Bravo}),
	)
}
