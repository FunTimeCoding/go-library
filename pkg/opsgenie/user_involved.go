package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/opsgenie/override"
	"github.com/funtimecoding/go-library/pkg/opsgenie/rotation"
)

func UserInvolved(
	user string,
	o []*override.Override,
	r []*rotation.Rotation,
) bool {
	for _, element := range o {
		if override.ContainsUser(element, user) {
			return true
		}
	}

	for _, element := range r {
		if rotation.ContainsUser(element, user) {
			return true
		}
	}

	return false
}
