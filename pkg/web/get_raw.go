package web

import "net/http"

func GetRaw(
	c *http.Client,
	locator string,
) *http.Response {
	return Send(c, NewGet(locator))
}
