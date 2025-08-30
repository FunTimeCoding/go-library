package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/data_source"
)

func (c *Client) DataSources() []*data_source.Source {
	result, _, e := c.client.CoreAPI.CoreDataSourcesList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return data_source.NewSlice(result.Results)
}
