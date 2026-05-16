package habitica

import (
	"github.com/funtimecoding/go-library/pkg/habitica/statistic"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) Allocate(stat string) (*statistic.Statistic, error) {
	var result *statistic.Statistic

	return result, c.post(
		join.Empty("/user/allocate?stat=", stat),
		nil,
		&result,
	)
}
