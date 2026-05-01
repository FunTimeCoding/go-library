package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"net"
)

func (c *Client) MustCreatePhysicalInterface(
	a net.HardwareAddr,
	description string,
	i *network.Interface,
) *physical_address.Address {
	result, e := c.CreatePhysicalInterface(a, description, i)
	errors.PanicOnError(e)

	return result
}
