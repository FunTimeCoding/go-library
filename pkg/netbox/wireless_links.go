package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/wireless_link"
)

func (c *Client) WirelessLinks() []*wireless_link.Link {
	result, r, e := c.client.WirelessAPI.WirelessWirelessLinksList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return wireless_link.NewSlice(result.Results)
}
