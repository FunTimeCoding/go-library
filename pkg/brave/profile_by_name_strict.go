package brave

import (
	"github.com/funtimecoding/go-library/pkg/brave/profile"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func ProfileByNameStrict(name string) *profile.Profile {
	p := ProfileByName(name)

	if p == nil {
		errors.NotFound(name)

		return &profile.Profile{}
	}

	return p
}
