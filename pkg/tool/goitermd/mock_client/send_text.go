package mock_client

import "fmt"

func (c *Client) SendText(
	identifier string,
	text string,
) error {
	s, okay := c.screens[identifier]

	if !okay {
		return fmt.Errorf("session %s not found", identifier)
	}

	s.Lines = append(s.Lines, text)
	c.screens[identifier] = s

	return nil
}
