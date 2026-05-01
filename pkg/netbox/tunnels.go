package netbox

import "github.com/funtimecoding/go-library/pkg/netbox/tunnel"

func (c *Client) Tunnels() ([]*tunnel.Tunnel, error) {
	result, _, e := c.client.VpnAPI.VpnTunnelsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return tunnel.NewSlice(result.Results), nil
}
