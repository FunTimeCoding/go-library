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
	t netbox.InterfaceTypeValue,
	h net.HardwareAddr,
) *network.Interface {
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
