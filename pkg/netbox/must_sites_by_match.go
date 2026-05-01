package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
)

func (c *Client) MustSitesByMatch(m string) []*site.Site {
	result, e := c.SitesByMatch(m)
	errors.PanicOnError(e)

	return result
}
