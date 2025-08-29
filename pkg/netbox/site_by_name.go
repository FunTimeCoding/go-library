package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
)

func (c *Client) SiteByName(n string) *site.Site {
	result := c.SitesByName(n)
	unexpected.Count(1, len(result))

	return result[0]
}
