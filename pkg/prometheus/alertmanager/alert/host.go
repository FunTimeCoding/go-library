package alert

import "github.com/funtimecoding/go-library/pkg/web/locator"

func (a *Alert) Host() string {
	return locator.RemoveDomain(a.FullHost())
}
