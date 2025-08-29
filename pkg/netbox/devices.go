package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) Devices() []*device.Device {
	result, _, e := c.client.DcimAPI.DcimDevicesList(
		c.context,
	).Limit(10000).Execute()
	errors.PanicOnError(e)

	return device.NewSlice(result.Results)
}
