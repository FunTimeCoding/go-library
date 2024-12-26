package request_context

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Context) SetLastLocation() {
	// Ignore if the referer header is set, favicon.ico is a common example.
	if _, okay := c.Header()[web.RefererHeader]; okay {
		return
	}

	c.SetCookie(web.LastLocationCookie, c.request.URL.Path)
}
