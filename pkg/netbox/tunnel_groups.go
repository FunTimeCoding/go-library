package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tunnel_group"
)

func (c *Client) TunnelGroups() []*tunnel_group.Group {
	result, r, e := c.client.VpnAPI.VpnTunnelGroupsList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return tunnel_group.NewSlice(result.Results)
}
