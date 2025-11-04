package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) UpdatePhysicalAddress(
	a *physical_address.Address,
	i *network.Interface,
) *physical_address.Address {
	q := netbox.NewMACAddressRequest(a.Name)
	q.SetAssignedObjectType(constant.InterfaceAddress)
	q.SetAssignedObjectId(int64(i.Identifier))
	result, r, e := c.client.DcimAPI.DcimMacAddressesUpdate(
		c.context,
		a.Identifier,
	).MACAddressRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return physical_address.New(result)
}
