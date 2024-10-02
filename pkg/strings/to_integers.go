package strings

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"strconv"
)

func ToIntegers(s []string) []int {
	result := make([]int, len(s))

	for i, element := range s {
		number, e := strconv.Atoi(element)
		errors.PanicOnError(e)
		result[i] = number
	}

	return result
}
