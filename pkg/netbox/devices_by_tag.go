package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) DevicesByTag(s string) []*device.Device {
	result, r, e := c.client.DcimAPI.DcimDevicesList(
		c.context,
	).Limit(1000).Tag([]string{s}).Execute()
	errors.PanicOnWebError(r, e)

	return device.NewSlice(result.Results)
}
