package confluence

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"

func (c *Client) SpaceByName(name string) *space.Space {
	for _, s := range c.Spaces() {
		if s.Name == name {
			return s
		}
	}

	return nil
}
