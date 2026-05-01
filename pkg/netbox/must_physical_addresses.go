package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
)

func (c *Client) MustPhysicalAddresses() []*physical_address.Address {
	result, e := c.PhysicalAddresses()
	errors.PanicOnError(e)

	return result
}
