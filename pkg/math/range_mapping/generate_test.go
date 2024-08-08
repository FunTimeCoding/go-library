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
			{Range: ranges.Range{L: 0, R: 0.625}, Value: "Alfa"},
			{Range: ranges.Range{L: 0.625, R: 1.25}, Value: "Bravo"},
			{Range: ranges.Range{L: 1.25, R: 1.875}, Value: "Charlie"},
			{Range: ranges.Range{L: 1.875, R: 2.5}, Value: "Delta"},
		},
		Generate(
			[]float64{0.5, 1.5, 2.5},
			[]string{
				strings.Alfa,
				strings.Bravo,
				strings.Charlie,
				strings.Delta,
			},
		),
	)
}
