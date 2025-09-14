package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/url"
)

func DecodeLocator(s string) string {
	result, e := url.QueryUnescape(s)
	errors.PanicOnError(e)

	return result
}
