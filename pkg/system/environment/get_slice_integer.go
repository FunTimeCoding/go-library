package environment

import "github.com/funtimecoding/go-library/pkg/strings"

func GetSliceInteger(name string) []int {
	return strings.ToIntegers(GetSlice(name))
}
