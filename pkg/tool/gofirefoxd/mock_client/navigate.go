package mock_client

import "fmt"

func (c *Client) Navigate(
	identifier int,
	l string,
) error {
	for i, t := range c.tabs {
		if t.Identifier == identifier {
			c.tabs[i].Locator = l

			return nil
		}
	}

	return fmt.Errorf("tab %d not found", identifier)
}
