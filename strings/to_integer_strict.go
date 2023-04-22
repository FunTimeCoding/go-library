package strings

import (
	"strconv"
	"strings"
)

func ToIntegerStrict(number string) int {
	parsed, fail := strconv.ParseInt(
		strings.TrimSpace(number),
		10,
		32,
	)

	if fail != nil {
		panic(fail)
	}

	return int(parsed)
}
