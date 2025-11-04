package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/prefix"
)

func (c *Client) Prefixes() []*prefix.Prefix {
	if len(c.cache.Prefixes) != 0 {
		return c.cache.Prefixes
	}

	result, r, e := c.client.IpamAPI.IpamPrefixesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)
	c.cache.Prefixes = prefix.NewSlice(result.Results)

	return c.cache.Prefixes
}
