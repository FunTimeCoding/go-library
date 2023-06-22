package web

import (
	"github.com/funtimecoding/go-library/errors"
	"net/http"
)

func Get(
	c *http.Client,
	locator string,
) *http.Response {
	request, e := http.NewRequest(
		GetMethod,
		locator,
		nil,
	)
	errors.PanicOnError(e)
	response, doFail := c.Do(request)
	errors.PanicOnError(doFail)

	return response
}
