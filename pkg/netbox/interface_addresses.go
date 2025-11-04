package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
)

func (c *Client) InterfaceAddresses(interfaceIdentifier int32) []*internet_address.Address {
	result, r, e := c.client.IpamAPI.IpamIpAddressesList(
		c.context,
	).Limit(constant.PageLimit).InterfaceId(
		[]int32{interfaceIdentifier},
	).Execute()
	errors.PanicOnWebError(r, e)

	return internet_address.NewSlice(result.Results)
}
