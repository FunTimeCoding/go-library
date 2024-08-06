package range_mapping

import (
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
	n := len(values)
	result := []*Mapping{{From: 0, To: f[0], Value: values[0]}}

	for i := 1; i < n; i++ {
		from := f[i-1]
		to := f[i]

		if i == n-1 {
			to = f[len(f)-1]
		}

		result = append(
			result,
			&Mapping{From: from, To: to, Value: values[i]},
		)
	}

	return result
}
