package habitica

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) UserStats() response.Stats {
	return c.user().Stats
}
