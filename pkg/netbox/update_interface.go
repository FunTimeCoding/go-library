package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

// UpdateInterface Update existing interface by name and assign MAC address to it
// If MAC address does not exist, it will be created
func (c *Client) UpdateInterface(
	d *device.Device,
	name string,
	t netbox.InterfaceTypeValue,
	h net.HardwareAddr,
) *network.Interface {
	p := c.PhysicalAddress(h)

	if p == nil {
		p = c.CreatePhysical(h, "")
	}

	i := c.DeviceInterfaceByNameStrict(d, name)
	var assigned bool

	if p.ObjectType == constant.InterfaceAddress {
		if int32(p.ObjectIdentifier) == i.Identifier {
			assigned = true
		}
	}

	if !assigned {
		c.AssignInterfaceToPhysical(p, i)
	}

	v := netbox.NewBriefDeviceRequest()
	v.SetName(d.Name)
	r := netbox.NewWritableInterfaceRequest(
		netbox.BriefDeviceRequestAsBriefInterfaceRequestDevice(v),
		name,
		t,
	)
	r.SetPrimaryMacAddress(
		netbox.BriefMACAddressRequestAsInterfaceRequestPrimaryMacAddress(
			netbox.NewBriefMACAddressRequest(h.String()),
		),
	)
	result, _, e := c.client.DcimAPI.DcimInterfacesUpdate(
		c.context,
		c.DeviceInterfaceByNameStrict(d, name).Identifier,
	).WritableInterfaceRequest(*r).Execute()
	errors.PanicOnError(e)

	return network.New(result)
}
