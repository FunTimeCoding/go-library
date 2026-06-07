package environment

import "github.com/funtimecoding/go-library/pkg/strings"

func SliceInteger64(name string) []int64 {
	return strings.MustToIntegers64(Slice(name))
}
