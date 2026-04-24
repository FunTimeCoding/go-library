package integers

import "math"

func FromUnsigned64(i uint64) int {
	if i > math.MaxInt {
		panic("out of int range")
	}

	return int(i)
}
