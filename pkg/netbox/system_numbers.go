package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/system_number"
)

func (c *Client) SystemNumbers() []*system_number.Number {
	result, r, e := c.client.IpamAPI.IpamAsnsList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return system_number.NewSlice(result.Results)
}
