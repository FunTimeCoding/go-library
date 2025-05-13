package web

import "github.com/funtimecoding/go-library/pkg/strings"

func HostPortFromLocatorSplit(s string) (string, int) {
	u := ParseLocator(s)

	return u.Hostname(), strings.ToIntegerStrict(u.Port())
}
