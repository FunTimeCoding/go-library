package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/custom_link"
)

func (c *Client) CustomLinks() []*custom_link.Link {
	result, _, e := c.client.ExtrasAPI.ExtrasCustomLinksList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return custom_link.NewSlice(result.Results)
}
