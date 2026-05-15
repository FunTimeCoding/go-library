package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) MustTags() []response.Tag {
	return c.tags
}
