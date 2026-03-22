package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"
)

func (c *Client) SpaceByNameStrict(name string) *space.Space {
	result := c.SpaceByName(name)

	if result == nil {
		panic(fmt.Sprintf("space not found: %s", name))
	}

	return result
}
