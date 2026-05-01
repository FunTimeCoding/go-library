package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/wireless_network"
)

func (c *Client) MustWirelessNetworks() []*wireless_network.Network {
	result, e := c.WirelessNetworks()
	errors.PanicOnError(e)

	return result
}
