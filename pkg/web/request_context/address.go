package request_context

import (
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Context) Address() string {
	if v, okay := c.Header()[web.RealAddressHeader]; okay {
		return v[0]
	}

	return network.SplitHost(c.request.RemoteAddr)
}
