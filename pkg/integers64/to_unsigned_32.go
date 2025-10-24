package integers64

import "math"

func ToUnsigned32(i int64) uint32 {
	if i < 0 || i > math.MaxUint32 {
		panic("out of uint32 range")
	}

	return uint32(i)
}
