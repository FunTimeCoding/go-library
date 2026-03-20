package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/configuration_context"
)

func (c *Client) ConfigurationContexts() []*configuration_context.Context {
	result, r, e := c.client.ExtrasAPI.ExtrasConfigContextsList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return configuration_context.NewSlice(result.Results)
}
