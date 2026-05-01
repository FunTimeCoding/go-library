package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
)

func (c *Client) MustUpdatePhysicalAddress(
	a *physical_address.Address,
	i *network.Interface,
) *physical_address.Address {
	result, e := c.UpdatePhysicalAddress(a, i)
	errors.PanicOnError(e)

	return result
}
