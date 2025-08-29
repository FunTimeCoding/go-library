package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) InternetAddressRanges() []netbox.IPRange {
	result, _, e := c.client.IpamAPI.IpamIpRangesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return result.Results
}
