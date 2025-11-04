package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
)

func (c *Client) Sites() []*site.Site {
	if len(c.cache.Sites) != 0 {
		return c.cache.Sites
	}

	result, r, e := c.client.DcimAPI.DcimSitesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)
	c.cache.Sites = site.NewSlice(result.Results)

	return c.cache.Sites
}
