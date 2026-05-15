package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) AddTag(tag response.Tag) {
	c.tags = append(c.tags, tag)
}
