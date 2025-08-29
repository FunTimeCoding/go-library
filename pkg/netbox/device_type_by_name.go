package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_type"
)

func (c *Client) DeviceTypeByName(n string) *device_type.Type {
	for _, element := range c.DeviceTypes() {
		if element.Model == n {
			return element
		}
	}

	errors.NotFound(n)

	return nil
}
