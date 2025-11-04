package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/wireless_network_group"
)

func (c *Client) WirelessNetworkGroups() []*wireless_network_group.Group {
	result, r, e := c.client.WirelessAPI.WirelessWirelessLanGroupsList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return wireless_network_group.NewSlice(result.Results)
}
