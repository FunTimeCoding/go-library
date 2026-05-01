package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) UnsetSerial(device string) (*device.Device, error) {
	d, e := c.DeviceByNameStrict(device)

	if e != nil {
		return nil, e
	}

	if c.verbose {
		fmt.Printf("set serial device: %+v\n", d)
		fmt.Printf("set serial raw device: %+v\n", d.Raw)
	}

	w := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	w.SetName(d.Name)
	w.SetSerial("")

	return c.updateDevice(d, w)
}
