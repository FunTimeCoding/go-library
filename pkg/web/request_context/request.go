package request_context

import "net/http"

func (c *Context) Request() *http.Request {
	return c.request
}
