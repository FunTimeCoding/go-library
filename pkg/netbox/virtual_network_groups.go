package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_network_group"
)

func (c *Client) VirtualNetworkGroups() []*virtual_network_group.Group {
	result, r, e := c.client.IpamAPI.IpamVlanGroupsList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return virtual_network_group.NewSlice(result.Results)
}
