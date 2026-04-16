package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateSite(name string) *site.Site {
	q := upstream.NewWritableSiteRequest(
		name,
		slug(name),
	)
	result, r, e := c.client.DcimAPI.DcimSitesCreate(
		c.context,
	).WritableSiteRequest(*q).Execute()
	errors.PanicOnWebError(r, e)
	c.cache.Sites = nil

	return site.New(result)
}
