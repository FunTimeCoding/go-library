package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/wireless_network_group"
)

func (c *Client) MustWirelessNetworkGroups() []*wireless_network_group.Group {
	result, e := c.WirelessNetworkGroups()
	errors.PanicOnError(e)

	return result
}
