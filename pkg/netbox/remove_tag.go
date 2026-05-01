package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) RemoveTag(
	deviceName string,
	tag string,
) (*device.Device, error) {
	d, e := c.DeviceByNameStrict(deviceName)

	if e != nil {
		return nil, e
	}

	if c.verbose {
		fmt.Printf("remove tag device: %+v\n", d)
		fmt.Printf("remove tag raw device: %+v\n", d.Raw)
	}

	d.RemoveTag(tag)
	w := devicePatch(d)
	tags, f := c.tagsNestedRequest(d.Tags)

	if f != nil {
		return nil, f
	}

	w.SetTags(tags)

	return c.updateDevice(d, w)
}
