package request_context

import "github.com/funtimecoding/go-library/pkg/web/constant"

func (c *Context) SetLastLocation() {
	// Ignore if the referer header is set, favicon.ico is a common example.
	if _, okay := c.Header()[constant.RefererHeader]; okay {
		return
	}

	c.SetCookie(constant.LastLocationCookie, c.request.URL.Path)
}
