package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

// CreatePhysicalInterface Create MAC address and assign to existing interface
func (c *Client) CreatePhysicalInterface(
	a net.HardwareAddr,
	description string,
	i *network.Interface,
) *physical_address.Address {
	r := netbox.NewMACAddressRequest(a.String())

	if description != "" {
		r.SetDescription(description)
	}

	r.SetAssignedObjectType(constant.InterfaceAddress)
	r.SetAssignedObjectId(int64(i.Identifier))

	result, _, e := c.client.DcimAPI.DcimMacAddressesCreate(
		c.context,
	).MACAddressRequest(*r).Execute()
	errors.PanicOnError(e)

	return physical_address.New(result)
}
