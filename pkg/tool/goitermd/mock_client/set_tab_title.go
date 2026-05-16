package mock_client

import "fmt"

func (c *Client) SetTabTitle(
	tabIdentifier string,
	title string,
) error {
	for i, s := range c.sessions {
		if s.TabId == tabIdentifier {
			c.sessions[i].TabTitle = title

			return nil
		}
	}

	return fmt.Errorf("tab %s not found", tabIdentifier)
}
