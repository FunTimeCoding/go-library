package locator

import "github.com/funtimecoding/go-library/pkg/strings/split"

func RemoveDomain(s string) string {
	if HasDot(s) {
		return split.Dot(s)[0]
	}

	return s
}
