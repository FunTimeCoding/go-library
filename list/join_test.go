package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"github.com/funtimecoding/go-library/strings"
	"testing"
)

func TestJoin(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"}, Join(
			[][]string{
				{
					strings.Alfa,
				},
				{
					strings.Bravo,
				},
			},
		),
	)
}
