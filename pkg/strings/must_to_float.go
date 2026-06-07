package strings

import (
	"strconv"
	"strings"
)

func MustToFloat(s string) float64 {
	result, e := strconv.ParseFloat(strings.TrimSpace(s), 64)

	if e != nil {
		panic(e)
	}

	return result
}
