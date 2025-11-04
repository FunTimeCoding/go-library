package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateInterface(
	d *device.Device,
	name string,
	t netbox.InterfaceTypeValue,
) *network.Interface {
	v := netbox.NewBriefDeviceRequest()
	v.SetName(d.Name)
	q := netbox.NewWritableInterfaceRequest(
		netbox.BriefDeviceRequestAsBriefInterfaceRequestDevice(v),
		name,
		t,
	)
	result, r, e := c.client.DcimAPI.DcimInterfacesCreate(
		c.context,
	).WritableInterfaceRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return network.New(result)
}
