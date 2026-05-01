package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
)

func (c *Client) MustInternetAddresses() []*internet_address.Address {
	result, e := c.InternetAddresses()
	errors.PanicOnError(e)

	return result
}
