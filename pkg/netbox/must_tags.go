package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tag"
)

func (c *Client) MustTags() []*tag.Tag {
	result, e := c.Tags()
	errors.PanicOnError(e)

	return result
}
