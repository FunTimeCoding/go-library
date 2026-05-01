package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
)

func (c *Client) MustSitesByName(n string) []*site.Site {
	result, e := c.SitesByName(n)
	errors.PanicOnError(e)

	return result
}
