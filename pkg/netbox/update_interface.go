package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) UpdateInterface(
	d *device.Device,
	name string,
	interfaceType netbox.InterfaceTypeValue,
	h net.HardwareAddr,
) *network.Interface {
	i := c.DeviceInterfaceByName(d, name, true)
	r2 := netbox.NewBriefDeviceRequest()
	r2.SetName(d.Name)
	r := netbox.NewWritableInterfaceRequest(
		netbox.BriefDeviceRequestAsBriefInterfaceRequestDevice(r2),
		name,
		interfaceType,
	)
	r.SetType(interfaceType)
	r.SetName(name)
	r.SetPrimaryMacAddress(
		netbox.BriefMACAddressRequestAsInterfaceRequestPrimaryMacAddress(
			netbox.NewBriefMACAddressRequest(h.String()),
		),
	)
	result, _, e := c.client.DcimAPI.DcimInterfacesUpdate(
		c.context,
		i.Identifier,
	).WritableInterfaceRequest(*r).Execute()
	errors.PanicOnError(e)

	return network.New(result, c.interfaceTypes)
}
