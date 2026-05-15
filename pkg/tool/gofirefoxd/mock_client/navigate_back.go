package mock_client

import "fmt"

func (c *Client) NavigateBack(identifier int) error {
	for _, t := range c.tabs {
		if t.Identifier == identifier {
			return nil
		}
	}

	return fmt.Errorf("tab %d not found", identifier)
}
