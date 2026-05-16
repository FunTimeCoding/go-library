package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/tag"

func (c *Client) MustTags() []*tag.Tag {
	return c.tags
}
