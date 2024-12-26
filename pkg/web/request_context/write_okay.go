package request_context

import "net/http"

func (c *Context) WriteOkay(s string) {
	c.Write(http.StatusOK, s)
}
