package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tunnel"
)

func (c *Client) Tunnels() []*tunnel.Tunnel {
	result, r, e := c.client.VpnAPI.VpnTunnelsList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return tunnel.NewSlice(result.Results)
}
