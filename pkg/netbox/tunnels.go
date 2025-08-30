package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tunnel"
)

func (c *Client) Tunnels() []*tunnel.Tunnel {
	result, _, e := c.client.VpnAPI.VpnTunnelsList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return tunnel.NewSlice(result.Results)
}
