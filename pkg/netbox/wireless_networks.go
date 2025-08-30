package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/wireless_network"
)

func (c *Client) WirelessNetworks() []*wireless_network.Network {
	result, _, e := c.client.WirelessAPI.WirelessWirelessLansList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return wireless_network.NewSlice(result.Results)
}
