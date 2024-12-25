package request_context

import "net/http"

func (c *Context) Cookie(name string) *http.Cookie {
	result, e := c.request.Cookie(name)

	if e != nil {
		return nil
	}

	return result
}
