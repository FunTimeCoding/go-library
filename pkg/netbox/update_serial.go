package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) UpdateSerial(
	device string,
	serial string,
) (*device.Device, error) {
	d, e := c.DeviceByNameStrict(device)

	if e != nil {
		return nil, e
	}

	if c.verbose {
		fmt.Printf("set serial device: %+v\n", d)
		fmt.Printf("set serial raw device: %+v\n", d.Raw)
	}

	d.Serial = serial

	return c.updateDevice(d, devicePatch(d))
}
