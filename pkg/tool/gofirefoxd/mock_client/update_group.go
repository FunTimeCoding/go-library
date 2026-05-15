package mock_client

import "fmt"

func (c *Client) UpdateGroup(
	groupIdentifier int,
	title string,
	color string,
	collapsed *bool,
) error {
	g, okay := c.groups[groupIdentifier]

	if !okay {
		return fmt.Errorf("group %d not found", groupIdentifier)
	}

	if title != "" {
		g.title = title
	}

	if color != "" {
		g.color = color
	}

	if collapsed != nil {
		g.collapsed = *collapsed
	}

	c.groups[groupIdentifier] = g

	return nil
}
