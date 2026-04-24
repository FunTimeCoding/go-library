package strings

import (
	"strconv"
	"strings"
)

func ToInteger64Strict(s string) int64 {
	result, e := strconv.ParseInt(strings.TrimSpace(s), 10, 64)

	if e != nil {
		panic(e)
	}

	return result
}
