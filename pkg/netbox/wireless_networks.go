package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/wireless_network"
)

func (c *Client) WirelessNetworks() []*wireless_network.Network {
	result, r, e := c.client.WirelessAPI.WirelessWirelessLansList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return wireless_network.NewSlice(result.Results)
}
