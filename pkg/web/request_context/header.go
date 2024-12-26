package request_context

import "net/http"

func (c *Context) Header() http.Header {
	return c.request.Header
}
