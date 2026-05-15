package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) UserGear() (response.Gear, error) {
	return c.gear, nil
}
