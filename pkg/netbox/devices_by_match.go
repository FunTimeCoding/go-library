package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) DevicesByMatch(s string) []*device.Device {
	result, _, e := c.client.DcimAPI.DcimDevicesList(
		c.context,
	).NameIc([]string{s}).Execute()
	errors.PanicOnError(e)

	return device.NewSlice(result.Results)
}
