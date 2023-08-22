package strings

import (
	"strconv"
	"strings"
)

func ToIntegerStrict(s string) int {
	result, e := strconv.ParseInt(strings.TrimSpace(s), 10, 32)

	if e != nil {
		panic(e)
	}

	return int(result)
}
