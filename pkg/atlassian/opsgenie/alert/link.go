package alert

import "github.com/funtimecoding/go-library/pkg/web/locator"

func link(
	identifier string,
	webHost string,
) string {
	return locator.New(
		webHost,
	).Path("/alert/detail/%s/details", identifier).String()
}
