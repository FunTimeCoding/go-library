package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
)

func (c *Client) DeviceAddresses(device string) []*internet_address.Address {
	result, _, e := c.client.IpamAPI.IpamIpAddressesList(
		c.context,
	).Device([]string{device}).Execute()
	errors.PanicOnError(e)

	return internet_address.NewSlice(result.Results)
}
