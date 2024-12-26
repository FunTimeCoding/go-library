package request_context

import "github.com/funtimecoding/go-library/pkg/network"

func (c *Context) Address() string {
	return network.SplitHost(c.request.RemoteAddr)
}
