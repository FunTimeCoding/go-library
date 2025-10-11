package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/url"
	"strings"
)

func ParseLocator(s string) *url.URL {
	if !strings.HasPrefix(s, constant.SecurePrefix) &&
		!strings.HasPrefix(s, constant.InsecurePrefix) {
		panic("locator must have a scheme")
	}

	// If no scheme is provided, the host becomes the scheme, which is wrong.
	result, e := url.Parse(s)
	errors.PanicOnError(e)

	return result
}
