package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tunnel"
)

func (c *Client) MustTunnels() []*tunnel.Tunnel {
	result, e := c.Tunnels()
	errors.PanicOnError(e)

	return result
}
