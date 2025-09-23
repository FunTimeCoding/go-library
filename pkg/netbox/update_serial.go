package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) UpdateSerial(
	device string,
	serial string,
) *device.Device {
	d := c.DeviceByNameStrict(device)

	if c.verbose {
		fmt.Printf("set serial device: %+v\n", d)
		fmt.Printf("set serial raw device: %+v\n", d.Raw)
	}

	d.Serial = serial
	w := devicePatch(d)

	return c.updateDevice(d, w)
}
