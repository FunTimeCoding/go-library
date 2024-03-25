package web

import "net/http"

func Post(
	c *http.Client,
	locator string,
	body string,
) *http.Response {
	return Send(c, NewPost(locator, body))
}
