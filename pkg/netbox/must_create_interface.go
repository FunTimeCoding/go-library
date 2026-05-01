package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) MustCreateInterface(d *device.Device, name string, t netbox.InterfaceTypeValue) *network.Interface {
	result, e := c.CreateInterface(d, name, t)
	errors.PanicOnError(e)

	return result
}
