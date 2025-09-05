package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) CreateInterfacePhysical(
	d *device.Device,
	name string,
	t netbox.InterfaceTypeValue,
	h net.HardwareAddr,
) *network.Interface {
	// Create interface with MAC address as primary

	// If mac address is not assigned to device already, setting it as primary fails

	// If mac address does not exist, create it or fail?
	//  TODO: Start with AssignInterfaceToPhysical
	//   Then create interface without MAC, then assign MAC to interface and make primary

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
	result, _, e := c.client.DcimAPI.DcimInterfacesCreate(
		c.context,
	).WritableInterfaceRequest(*r).Execute()
	errors.PanicOnError(e)

	return network.New(result)
}
