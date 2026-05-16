package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/statistic"

func (c *Client) UserStats() (*statistic.Statistic, error) {
	return c.stats, nil
}
