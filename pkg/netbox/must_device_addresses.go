package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
)

func (c *Client) MustDeviceAddresses(device string) []*internet_address.Address {
	result, e := c.DeviceAddresses(device)
	errors.PanicOnError(e)

	return result
}
