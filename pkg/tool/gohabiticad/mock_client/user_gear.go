package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/gear"

func (c *Client) UserGear() (*gear.Gear, error) {
	return c.gear, nil
}
