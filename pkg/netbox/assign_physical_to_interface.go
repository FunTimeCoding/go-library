package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) AssignPhysicalToInterface(
	p *physical_address.Address,
	i *network.Interface,
) *physical_address.Address {
	r := netbox.NewMACAddressRequest(p.Name)

	if i == nil {
		r.SetAssignedObjectTypeNil()
		r.SetAssignedObjectIdNil()
	} else {
		r.SetAssignedObjectType(constant.InterfaceAddress)
		r.SetAssignedObjectId(int64(i.Identifier))
	}

	result, _, e := c.client.DcimAPI.DcimMacAddressesUpdate(
		c.context,
		p.Identifier,
	).MACAddressRequest(*r).Execute()
	errors.PanicOnError(e)

	return physical_address.New(result)
}
