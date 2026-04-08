package environment

import "github.com/funtimecoding/go-library/pkg/strings"

func RequiredInteger(name string) int {
	return strings.ToIntegerStrict(Required(name))
}
