package web

import "net/http"

func GetString(
	c *http.Client,
	locator string,
) string {
	return ReadString(Get(c, locator))
}
