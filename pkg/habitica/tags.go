package habitica

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) Tags() ([]response.Tag, error) {
	var result []response.Tag

	return result, c.get("/tags", &result)
}
