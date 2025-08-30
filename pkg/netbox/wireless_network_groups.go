package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/wireless_network_group"
)

func (c *Client) WirelessNetworkGroups() []*wireless_network_group.Group {
	result, _, e := c.client.WirelessAPI.WirelessWirelessLanGroupsList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return wireless_network_group.NewSlice(result.Results)
}
