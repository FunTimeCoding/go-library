package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestSecondHalf(t *testing.T) {
	assert.Any(
		t,
		[]string{"Charlie", "Delta"},
		SecondHalf(
			[]string{
				upper.Alfa,
				upper.Bravo,
				upper.Charlie,
				upper.Delta,
			},
		),
	)
}
