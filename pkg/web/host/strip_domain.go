package host

import "github.com/funtimecoding/go-library/pkg/strings/split"

func StripDomain(s string) string {
	if HasDot(s) {
		return split.Dot(s)[0]
	}

	return s
}
