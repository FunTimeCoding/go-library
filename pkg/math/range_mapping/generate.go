package range_mapping

import (
	"github.com/funtimecoding/go-library/pkg/math/ranges"
	"slices"
	"sort"
)

func Generate(
	f []float64,
	values []string,
) []*Mapping {
	if slices.Contains(f, 0) {
		f = append(f, 0)
	}

	sort.Float64s(f)
	maxFloat := f[len(f)-1]
	count := len(values)

	if count == 0 {
		return nil
	}

	mappings := maxFloat / float64(count)
	result := make([]*Mapping, 0)

	for i := 0; i < count; i++ {
		result = append(
			result,
			&Mapping{
				Range: ranges.Range{
					L: float64(i) * mappings,
					R: float64(i+1) * mappings,
				},
				Value: values[i],
			},
		)
	}

	return result
}
