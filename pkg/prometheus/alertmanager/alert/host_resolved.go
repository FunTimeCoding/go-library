package alert

import "github.com/funtimecoding/go-library/pkg/web/address"

func (a *Alert) HostResolved() bool {
	return !address.IsInternet(a.FullHost())
}
