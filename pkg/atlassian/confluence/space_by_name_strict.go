package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"
)

func (c *Client) SpaceByNameStrict(name string) (*space.Space, error) {
	result, e := c.SpaceByName(name)

	if e != nil {
		return nil, e
	}

	if result == nil {
		return nil, fmt.Errorf("space not found: %s", name)
	}

	return result, nil
}
