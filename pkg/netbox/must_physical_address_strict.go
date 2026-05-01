package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"net"
)

func (c *Client) MustPhysicalAddressStrict(a net.HardwareAddr) *physical_address.Address {
	result, e := c.PhysicalAddressStrict(a)
	errors.PanicOnError(e)

	return result
}
