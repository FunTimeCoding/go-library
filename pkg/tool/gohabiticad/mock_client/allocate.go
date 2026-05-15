package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) Allocate(_ string) (response.Stats, error) {
	return c.stats, nil
}
