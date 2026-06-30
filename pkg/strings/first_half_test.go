package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestFirstHalf(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		FirstHalf(
			[]string{
				upper.Alfa,
				upper.Bravo,
				upper.Charlie,
				upper.Delta,
			},
		),
	)
}
