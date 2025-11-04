package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) physicalAddressesOffset(offset int32) []netbox.MACAddress {
	result, r, e := c.client.DcimAPI.DcimMacAddressesList(
		c.context,
	).Limit(constant.PageLimit).Offset(offset).Execute()
	errors.PanicOnWebError(r, e)

	return result.Results
}
