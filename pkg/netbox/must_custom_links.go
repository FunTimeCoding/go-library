package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/custom_link"
)

func (c *Client) MustCustomLinks() []*custom_link.Link {
	result, e := c.CustomLinks()
	errors.PanicOnError(e)

	return result
}
