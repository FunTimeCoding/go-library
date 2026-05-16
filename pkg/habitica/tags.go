package habitica

import "github.com/funtimecoding/go-library/pkg/habitica/tag"

func (c *Client) Tags() ([]*tag.Tag, error) {
	var result []*tag.Tag

	return result, c.get("/tags", &result)
}
