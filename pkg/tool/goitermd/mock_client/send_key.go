package mock_client

import "fmt"

func (c *Client) SendKey(
	identifier string,
	_ string,
) error {
	if _, okay := c.screens[identifier]; !okay {
		return fmt.Errorf("session %s not found", identifier)
	}

	return nil
}
