package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

// CreateInterfacePhysical Create new interface and assign MAC address to it
// If MAC address does not exist, it will be created
func (c *Client) CreateInterfacePhysical(
	d *device.Device,
	name string,
	t netbox.InterfaceTypeValue,
	h net.HardwareAddr,
) *network.Interface {
	p := c.PhysicalAddress(h)

	if p == nil {
		p = c.CreatePhysical(h, "")
	}

	v := netbox.NewBriefDeviceRequest()
	v.SetName(d.Name)
	// skip setting as primary, needs to be assigned first; then assign
	c.AssignPhysicalToInterface(
		p,
		c.createInterfaceWriteable(
			netbox.NewWritableInterfaceRequest(
				netbox.BriefDeviceRequestAsBriefInterfaceRequestDevice(v),
				name,
				t,
			),
		),
	)

	// set as primary
	return c.UpdateInterface(d, name, t, h)
}
