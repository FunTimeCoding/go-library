package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) internetAddressesOffset(offset int32) []netbox.IPAddress {
	result, _, e := c.client.IpamAPI.IpamIpAddressesList(
		c.context,
	).Limit(constant.PageLimit).Offset(offset).Execute()
	errors.PanicOnError(e)

	return result.Results
}
