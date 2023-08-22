package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToLower(t *testing.T) {
	assert.Strings(
		t,
		[]string{"alfa", "bravo"},
		ToLower([]string{Alfa, Bravo}),
	)
}
