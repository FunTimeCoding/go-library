package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_network"
)

func (c *Client) MustVirtualNetworks() []*virtual_network.Network {
	result, e := c.VirtualNetworks()
	errors.PanicOnError(e)

	return result
}
