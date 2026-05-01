package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
)

func (c *Client) MustCreateAddress(interfaceIdentifier int32, address string) *internet_address.Address {
	result, e := c.CreateAddress(interfaceIdentifier, address)
	errors.PanicOnError(e)

	return result
}
