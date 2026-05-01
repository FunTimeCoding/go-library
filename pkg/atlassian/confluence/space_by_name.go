package confluence

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"

func (c *Client) SpaceByName(name string) (*space.Space, error) {
	spaces, e := c.Spaces()

	if e != nil {
		return nil, e
	}

	for _, s := range spaces {
		if s.Name == name {
			return s, nil
		}
	}

	return nil, nil
}
