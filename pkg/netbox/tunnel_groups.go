package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tunnel_group"
)

func (c *Client) TunnelGroups() []*tunnel_group.Group {
	result, _, e := c.client.VpnAPI.VpnTunnelGroupsList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return tunnel_group.NewSlice(result.Results)
}
