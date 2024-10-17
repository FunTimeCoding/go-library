package web

import "net/http"

func Patch(
	c *http.Client,
	locator string,
	body string,
) *http.Response {
	return Send(c, NewPatch(locator, body))
}
