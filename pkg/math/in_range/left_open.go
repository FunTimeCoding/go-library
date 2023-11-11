package in_range

import "github.com/funtimecoding/go-library/pkg/math/ranges"

func LeftOpen(
	value float64,
	interval ranges.Range,
) bool {
	if value > interval.L && value <= interval.R {
		return true
	}

	return false
}
