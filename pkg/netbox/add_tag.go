package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) AddTag(
	deviceName string,
	tag string,
) *device.Device {
	t := c.TagByName(tag)

	if t == nil {
		// Does not happen because TagByName panics if not found

		return nil
	}

	if c.verbose {
		fmt.Printf(
			"ADD tag: %s %s\n",
			t.Nested.GetName(),
			t.Nested.GetSlug(),
		)
	}

	d := c.DeviceByNameStrict(deviceName)

	if c.verbose {
		fmt.Printf("ADD device: %+v\n", d)
		fmt.Printf("ADD raw device: %+v\n", d.Raw)
	}

	d.AddTag(t.Name)
	w := devicePatch(d)
	w.SetTags(c.tagsNestedRequest(d.Tags))

	return c.updateDevice(d, w)
}
