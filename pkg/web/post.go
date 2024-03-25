package web

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func Post(
	c *http.Client,
	locator string,
	body string,
) *http.Response {
	r, e := http.NewRequest(
		PostMethod,
		locator,
		bytes.NewBuffer([]byte(body)),
	)
	errors.PanicOnError(e)

	return Send(c, r)
}
