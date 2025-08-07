package request_context

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/location"
)

func (c *Context) LastLocation() string {
	if s := c.Cookie(constant.LastLocationCookie); s != nil {
		return s.Value
	}

	return location.Root
}
