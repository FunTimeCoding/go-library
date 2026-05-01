package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
)

func (c *Client) MustInterfaceAddresses(interfaceIdentifier int32) []*internet_address.Address {
	result, e := c.InterfaceAddresses(interfaceIdentifier)
	errors.PanicOnError(e)

	return result
}
