package request_context

import "github.com/funtimecoding/go-library/pkg/web/constant"

func (c *Context) ClientAddress() string {
	if s := c.request.Header.Get(constant.ForwardedFor); s != "" {
		return s
	}

	if s := c.request.Header.Get(constant.RealAddress); s != "" {
		return s
	}

	return c.request.RemoteAddr
}
