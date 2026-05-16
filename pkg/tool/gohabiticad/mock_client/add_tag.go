package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/tag"

func (c *Client) AddTag(t *tag.Tag) {
	c.tags = append(c.tags, t)
}
