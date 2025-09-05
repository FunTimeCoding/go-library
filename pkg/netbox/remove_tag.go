package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) RemoveTag(
	deviceName string,
	tag string,
) *device.Device {
	d := c.DeviceByNameStrict(deviceName)

	if c.verbose {
		fmt.Printf("REMOVE device: %+v\n", d)
		fmt.Printf("REMOVE raw device: %+v\n", d.Raw)
	}

	d.RemoveTag(tag)
	w := devicePatch(d)
	w.SetTags(c.tagsNestedRequest(d.Tags))

	return c.updateDevice(d, w)
}
