package request_context

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Context) QueryList(k string) []string {
	return web.GetList(c.request, k)
}
