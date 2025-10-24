package integers

import "math"

func To32(x int) int32 {
	if x > math.MaxInt32 || x < math.MinInt32 {
		panic("out of int32 range")
	}

	return int32(x)
}
