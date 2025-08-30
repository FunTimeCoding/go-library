package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/system_number"
)

func (c *Client) SystemNumbers() []*system_number.Number {
	result, _, e := c.client.IpamAPI.IpamAsnsList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return system_number.NewSlice(result.Results)
}
