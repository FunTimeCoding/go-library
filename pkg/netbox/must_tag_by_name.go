package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tag"
)

func (c *Client) MustTagByName(n string) *tag.Tag {
	result, e := c.TagByName(n)
	errors.PanicOnError(e)

	return result
}
