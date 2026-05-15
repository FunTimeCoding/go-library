package mock_client

import "fmt"

func (c *Client) CloseTab(identifier int) error {
	for i, t := range c.tabs {
		if t.Identifier == identifier {
			c.tabs = append(c.tabs[:i], c.tabs[i+1:]...)

			return nil
		}
	}

	return fmt.Errorf("tab %d not found", identifier)
}
