package request_context

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Context) ClientAddress() string {
	return web.ClientAddress(c.request)
}
