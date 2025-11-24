package environment

import "github.com/funtimecoding/go-library/pkg/strings"

func SliceInteger(name string) []int {
	return strings.ToIntegersStrict(Slice(name))
}
