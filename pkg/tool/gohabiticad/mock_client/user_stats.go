package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) UserStats() (response.Stats, error) {
	return c.stats, nil
}
