package web

import (
	"io"
	"net/http"
)

func PostBytes(
	c *http.Client,
	locator string,
	body io.Reader,
) *http.Response {
	return Send(c, NewPostBytes(locator, body))
}
