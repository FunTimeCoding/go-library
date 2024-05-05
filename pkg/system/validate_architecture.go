package system

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"slices"
)

func ValidateArchitecture(s string) {
	if !slices.Contains(Architectures, s) {
		unexpected.String(s)
	}
}
