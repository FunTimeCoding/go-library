package locator

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/split"
)

func RemoveSubdomain(s string) string {
	if IsSubdomain(s) {
		return join.Dot(split.Dot(s)[1:])
	}

	return s
}
