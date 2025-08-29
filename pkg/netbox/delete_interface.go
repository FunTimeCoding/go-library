package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) DeleteInterface(
	d *device.Device,
	name string,
) {
	i := c.DeviceInterfaceByName(d, name, true)
	result, e := c.client.DcimAPI.DcimInterfacesDestroy(
		c.context,
		i.Identifier,
	).Execute()
	fmt.Printf("Delete interface result: %+v\n", result)
	errors.PanicOnError(e)
}
