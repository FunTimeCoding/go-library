package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func Send(
	c *http.Client,
	r *http.Request,
) *http.Response {
	result, e := c.Do(r)
	errors.PanicOnError(e)

	return result
}
