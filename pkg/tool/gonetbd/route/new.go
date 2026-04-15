package route

import "github.com/funtimecoding/go-library/pkg/netbox"

func New(c *netbox.Client) *Router {
	return &Router{client: c}
}
