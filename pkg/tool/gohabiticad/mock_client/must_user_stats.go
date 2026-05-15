package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) MustUserStats() response.Stats {
	return c.stats
}
