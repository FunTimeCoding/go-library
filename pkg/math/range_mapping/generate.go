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
	groupSize := len(f) / n
	var result []*Mapping

	for i := 0; i < n; i++ {
		from := i * groupSize
		to := (i+1)*groupSize - 1

		if i == n-1 {
			to = len(f) - 1
		}

		result = append(
			result,
			&Mapping{
				From:  f[from],
				To:    f[to],
				Value: values[i],
			},
		)
	}

	return result
}
