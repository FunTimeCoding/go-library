package alert

import "github.com/funtimecoding/go-library/pkg/web/host"

func (a *Alert) Host() string {
	return host.StripDomain(a.FullHost())
}
