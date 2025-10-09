package confluence

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"

func (c *Client) SpaceByNameStrict(name string) *space.Space {
	result := c.SpaceByName(name)

	if result == nil {
		panic("space not found: " + name)
	}

	return result
}
