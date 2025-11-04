package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) Devices() []*device.Device {
	result, r, e := c.client.DcimAPI.DcimDevicesList(
		c.context,
	).Limit(10000).Execute()
	errors.PanicOnWebError(r, e)

	return device.NewSlice(result.Results)
}
