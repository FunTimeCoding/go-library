package floats

import "github.com/funtimecoding/go-library/math"

func ToStringRounded(number float64) string {
	return ToString(math.Round(number, 1))
}
