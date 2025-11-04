package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/rack"
)

func (c *Client) RacksByName(n string) []*rack.Rack {
	result, r, e := c.client.DcimAPI.DcimRacksList(
		c.context,
	).Name([]string{n}).Execute()
	errors.PanicOnWebError(r, e)

	return rack.NewSlice(result.Results)
}
