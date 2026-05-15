package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) Tags() ([]response.Tag, error) {
	return c.tags, nil
}
