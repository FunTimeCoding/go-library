package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/source"
)

func (c *Client) Sources() []*source.Source {
	result, r, e := c.client.CoreAPI.CoreDataSourcesList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return source.NewSlice(result.Results)
}
