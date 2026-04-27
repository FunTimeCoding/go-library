package habitica

import "github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica/response"

func (c *Client) UserStats() response.Stats {
	return c.user().Stats
}
