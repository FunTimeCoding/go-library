package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/url"
)

func ParseLocator(s string) *url.URL {
	result, e := url.Parse(s)
	errors.PanicOnError(e)

	return result
}
