package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
)

func (c *Client) MustCreateSite(name string) *site.Site {
	result, e := c.CreateSite(name)
	errors.PanicOnError(e)

	return result
}
