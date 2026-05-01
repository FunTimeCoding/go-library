package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDeleteInterfaceByName(d *device.Device, name string) {
	errors.PanicOnError(c.DeleteInterfaceByName(d, name))
}
