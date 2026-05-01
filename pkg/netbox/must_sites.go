package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
)

func (c *Client) MustSites() []*site.Site {
	result, e := c.Sites()
	errors.PanicOnError(e)

	return result
}
