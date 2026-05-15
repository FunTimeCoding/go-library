package mock_client

import "fmt"

func (c *Client) UngroupTab(identifier int) error {
	for i, t := range c.tabs {
		if t.Identifier == identifier {
			c.tabs[i].GroupId = 0

			return nil
		}
	}

	return fmt.Errorf("tab %d not found", identifier)
}
