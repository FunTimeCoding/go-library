package range_mapping

import (
	"github.com/funtimecoding/go-library/pkg/math/ranges"
	"sort"
)

func Generate(
	f []float64,
	values []string,
) []*Mapping {
	sort.Float64s(f)
	result := make([]*Mapping, 0)
	rangeCount := len(values)

	if rangeCount == 0 {
		return result
	}

	maximum := f[len(f)-1]
	perRange := maximum / float64(rangeCount)

	for i := 0; i < rangeCount; i++ {
		m := New()
		m.Range = ranges.Range{
			L: float64(i) * perRange,
			R: float64(i+1) * perRange,
		}
		m.Value = values[i]
		result = append(result, m)
	}

	return result
}
