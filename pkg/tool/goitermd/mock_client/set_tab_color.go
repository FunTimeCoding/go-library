package mock_client

import "fmt"

func (c *Client) SetTabColor(
	sessionIdentifier string,
	_ int,
	_ int,
	_ int,
) error {
	if _, okay := c.screens[sessionIdentifier]; !okay {
		return fmt.Errorf("session %s not found", sessionIdentifier)
	}

	return nil
}
