package strings

import (
	"github.com/funtimecoding/go-library/errors"
	"strconv"
)

func ToIntegers(input []string) []int {
	result := make([]int, len(input))

	for i, element := range input {
		number, e := strconv.Atoi(element)
		errors.PanicOnError(e)
		result[i] = number
	}

	return result
}
