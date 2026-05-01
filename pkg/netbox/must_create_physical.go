package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"net"
)

func (c *Client) MustCreatePhysical(a net.HardwareAddr, description string) *physical_address.Address {
	result, e := c.CreatePhysical(a, description)
	errors.PanicOnError(e)

	return result
}
