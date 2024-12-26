package request_context

import "net/http"

func (c *Context) Redirect(target string) {
	http.Redirect(
		c.writer,
		c.request,
		target,
		http.StatusFound,
	)
}
