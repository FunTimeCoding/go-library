package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/config_context"
)

func (c *Client) ConfigContexts() []*config_context.Context {
	result, _, e := c.client.ExtrasAPI.ExtrasConfigContextsList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return config_context.NewSlice(result.Results)
}
