package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/rack"
)

func (c *Client) RacksByName(n string) []*rack.Rack {
	result, _, e := c.client.DcimAPI.DcimRacksList(
		c.context,
	).Name([]string{n}).Execute()
	errors.PanicOnError(e)

	return rack.NewSlice(result.Results)
}
