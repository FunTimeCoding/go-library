package request_context

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Context) UnsetCookie(k string) {
	web.UnsetCookie(c.writer, k)
}
