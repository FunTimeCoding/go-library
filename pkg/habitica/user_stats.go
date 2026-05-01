package habitica

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) UserStats() (response.Stats, error) {
	result, e := c.user()

	return result.Stats, e
}
