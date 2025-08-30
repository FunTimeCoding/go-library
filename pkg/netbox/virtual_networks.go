package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_network"
)

func (c *Client) VirtualNetworks() []*virtual_network.Network {
	result, _, e := c.client.IpamAPI.IpamVlansList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return virtual_network.NewSlice(result.Results)
}
