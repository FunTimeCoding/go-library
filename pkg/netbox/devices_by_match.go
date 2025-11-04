package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) DevicesByMatch(s string) []*device.Device {
	result, r, e := c.client.DcimAPI.DcimDevicesList(
		c.context,
	).NameIc([]string{s}).Execute()
	errors.PanicOnWebError(r, e)

	return device.NewSlice(result.Results)
}
