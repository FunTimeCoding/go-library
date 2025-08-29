package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
)

func (c *Client) SitesByMatch(m string) []*site.Site {
	result, _, e := c.client.DcimAPI.DcimSitesList(
		c.context,
	).NameIc([]string{m}).Execute()
	errors.PanicOnError(e)

	return site.NewSlice(result.Results)
}
