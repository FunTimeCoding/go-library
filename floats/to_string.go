package floats

import "strconv"

func ToString(number float64) string {
	return strconv.FormatFloat(number, 'f', -1, 32)
}
