package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) DevicesByRack(i int32) []*device.Device {
	result, _, e := c.client.DcimAPI.DcimDevicesList(
		c.context,
	).RackId([]int32{i}).Execute()
	errors.PanicOnError(e)

	return device.NewSlice(result.Results)
}
