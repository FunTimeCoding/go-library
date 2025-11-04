package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
)

func (c *Client) SitesByName(n string) []*site.Site {
	result, r, e := c.client.DcimAPI.DcimSitesList(
		c.context,
	).Name([]string{n}).Execute()
	errors.PanicOnWebError(r, e)

	return site.NewSlice(result.Results)
}
