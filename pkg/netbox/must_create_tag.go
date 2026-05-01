package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tag"
)

func (c *Client) MustCreateTag(name string) *tag.Tag {
	result, e := c.CreateTag(name)
	errors.PanicOnError(e)

	return result
}
