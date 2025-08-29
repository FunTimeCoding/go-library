package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) CreateInterface(
	d *device.Device,
	name string,
	n netbox.InterfaceTypeValue,
	h net.HardwareAddr,
) *network.Interface {
	r2 := netbox.NewBriefDeviceRequest()
	r2.SetName(d.Name)
	r := netbox.NewWritableInterfaceRequest(
		netbox.BriefDeviceRequestAsBriefInterfaceRequestDevice(r2),
		name,
		n,
	)
	r.SetPrimaryMacAddress(
		netbox.BriefMACAddressRequestAsInterfaceRequestPrimaryMacAddress(
			netbox.NewBriefMACAddressRequest(h.String()),
		),
	)
	result, _, e := c.client.DcimAPI.DcimInterfacesCreate(
		c.context,
	).WritableInterfaceRequest(*r).Execute()
	errors.PanicOnError(e)

	return network.New(result, c.interfaceTypes)
}
