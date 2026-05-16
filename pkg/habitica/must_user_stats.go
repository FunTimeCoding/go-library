package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/statistic"
)

func (c *Client) MustUserStats() *statistic.Statistic {
	result, e := c.UserStats()
	errors.PanicOnError(e)

	return result
}
