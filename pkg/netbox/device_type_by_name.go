package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_type"
)

func (c *Client) DeviceTypeByName(n string) *device_type.Type {
	for _, t := range c.DeviceTypes() {
		if t.Model == n {
			return t
		}
	}

	errors.NotFound(n)

	return nil
}
