package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/tag"
)

func (c *Client) Tags() []*tag.Tag {
	if len(c.cache.Tags) != 0 {
		return c.cache.Tags
	}

	result, r, e := c.client.ExtrasAPI.ExtrasTagsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)
	c.cache.Tags = tag.NewSlice(result.Results)

	return c.cache.Tags
}
