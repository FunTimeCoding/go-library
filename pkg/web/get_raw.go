package web

import "net/http"

func Get(
	c *http.Client,
	locator string,
) *http.Response {
	return Send(c, NewGet(locator))
}
