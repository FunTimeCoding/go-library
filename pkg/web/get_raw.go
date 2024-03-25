package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func Get(
	c *http.Client,
	locator string,
) *http.Response {
	r, e := http.NewRequest(GetMethod, locator, nil)
	errors.PanicOnError(e)

	return Send(c, r)
}
