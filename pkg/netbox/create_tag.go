package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tag"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateTag(name string) *tag.Tag {
	q := upstream.NewTagRequest(name, slug(name))
	result, r, e := c.client.ExtrasAPI.ExtrasTagsCreate(
		c.context,
	).TagRequest(*q).Execute()
	errors.PanicOnWebError(r, e)
	c.cache.Tags = nil

	return tag.New(result)
}
