package request_context

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (c *Context) SetCookie(k string, v string) *http.Cookie {
	return web.SetCookie(c.writer, k, v)
}
