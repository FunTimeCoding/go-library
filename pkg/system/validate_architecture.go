package system

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"slices"
)

func ValidateArchitecture(s string) {
	if !slices.Contains(constant.Architectures, s) {
		unexpected.String(s)
	}
}
