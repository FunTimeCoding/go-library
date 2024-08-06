package range_mapping

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	assert.Any(
		t,
		[]*Mapping{
			{From: 0, To: 0.5, Value: "Alfa"},
			{From: 0.5, To: 1.5, Value: "Bravo"},
			{From: 1.5, To: 2.5, Value: "Charlie"},
		},
		Generate(
			[]float64{0.5, 1.5, 2.5},
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
		),
	)
}
