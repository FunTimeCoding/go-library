package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFirstHalf(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		FirstHalf([]string{Alfa, Bravo, Charlie, Delta}),
	)
}
