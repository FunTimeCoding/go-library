package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) AssignInterfaceToPhysical(
	p *physical_address.Address,
	i *network.Interface,
) *netbox.MACAddress {
	r := netbox.NewMACAddressRequest(p.Name)

	if i == nil {
		// TODO: Is this how to unset?
		r.UnsetAssignedObjectId()
		r.UnsetAssignedObjectType()
	} else {
		r.SetAssignedObjectType(constant.InterfaceAddress)
		r.SetAssignedObjectId(int64(i.Identifier))
	}

	result, _, e := c.client.DcimAPI.DcimMacAddressesUpdate(
		c.context,
		p.Identifier,
	).MACAddressRequest(*r).Execute()
	errors.PanicOnError(e)

	return result
}
