package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/config_context"
)

func (c *Client) ConfigContexts() []*config_context.Context {
	result, r, e := c.client.ExtrasAPI.ExtrasConfigContextsList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return config_context.NewSlice(result.Results)
}
