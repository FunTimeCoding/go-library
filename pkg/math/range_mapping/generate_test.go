package range_mapping

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/math/ranges"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	assert.Any(
		t,
		[]*Mapping{
			{Range: ranges.Range{L: 0, R: 0.5}, Value: "Alfa"},
			{Range: ranges.Range{L: 0.5, R: 1.5}, Value: "Bravo"},
			{Range: ranges.Range{L: 1.5, R: 2.5}, Value: "Charlie"},
		},
		Generate(
			[]float64{0.5, 1.5, 2.5},
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
		),
	)
}
