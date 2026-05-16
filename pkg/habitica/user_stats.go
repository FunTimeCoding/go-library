package habitica

import "github.com/funtimecoding/go-library/pkg/habitica/statistic"

func (c *Client) UserStats() (*statistic.Statistic, error) {
	result, e := c.user()

	return result.Stats, e
}
