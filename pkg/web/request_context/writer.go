package request_context

import "net/http"

func (c *Context) Writer() http.ResponseWriter {
	return c.writer
}
