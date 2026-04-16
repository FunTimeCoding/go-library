package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateVirtualAddress(
	interfaceIdentifier int32,
	address string,
) *internet_address.Address {
	q := upstream.NewWritableIPAddressRequest(address)
	q.SetAssignedObjectType("virtualization.vminterface")
	q.SetAssignedObjectId(int64(interfaceIdentifier))
	result, r, e := c.client.IpamAPI.IpamIpAddressesCreate(
		c.context,
	).WritableIPAddressRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return internet_address.New(result)
}
