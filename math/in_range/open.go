package in_range

import "github.com/funtimecoding/go-library/math/ranges"

func Open(
	value float64,
	interval ranges.Range,
) bool {
	if value > interval.L && value < interval.R {
		return true
	}

	return false
}
