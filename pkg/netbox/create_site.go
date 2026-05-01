package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/site"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateSite(name string) (*site.Site, error) {
	q := upstream.NewWritableSiteRequest(
		name,
		slug(name),
	)
	result, _, e := c.client.DcimAPI.DcimSitesCreate(
		c.context,
	).WritableSiteRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.Sites = nil

	return site.New(result), nil
}
