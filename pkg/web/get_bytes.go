package web

import "net/http"

func GetBytes(
	c *http.Client,
	locator string,
) []byte {
	return ReadBytes(GetRaw(c, locator))
}
