package habitica

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) UserGear() (response.Gear, error) {
	result, e := c.user()

	return result.Items.Gear, e
}
