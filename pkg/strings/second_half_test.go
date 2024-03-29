package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSecondHalf(t *testing.T) {
	assert.Any(
		t,
		[]string{"Charlie", "Delta"},
		SecondHalf([]string{Alfa, Bravo, Charlie, Delta}),
	)
}
