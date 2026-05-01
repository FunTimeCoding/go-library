package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
)

func (c *Client) MustSiteByName(n string) *site.Site {
	result, e := c.SiteByName(n)
	errors.PanicOnError(e)

	return result
}
