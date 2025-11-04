package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) CreateInternet(
	objectType string,
	objectIdentifier int64,
	i net.IP,
	m net.IPMask,
) *netbox.IPAddress {
	address := AddressMask(i, m)
	fmt.Printf("Address: %s\n", address)
	q := netbox.NewWritableIPAddressRequest(address)
	q.SetAssignedObjectType(objectType)
	q.SetAssignedObjectId(objectIdentifier)
	result, r, e := c.client.IpamAPI.IpamIpAddressesCreate(
		c.context,
	).WritableIPAddressRequest(*q).Execute()
	fmt.Printf("Create address result: %+v\n", result)
	errors.PanicOnWebError(r, e)

	return result
}
