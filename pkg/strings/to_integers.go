package strings

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"strconv"
)

func ToIntegers(s []string) []int {
	result := make([]int, len(s))

	for i, l := range s {
		number, e := strconv.Atoi(l)
		errors.PanicOnError(e)
		result[i] = number
	}

	return result
}
