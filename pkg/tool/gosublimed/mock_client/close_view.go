package mock_client

import "fmt"

func (c *Client) CloseView(identifier int) error {
	for i, v := range c.views {
		if v.Identifier == identifier {
			c.views = append(c.views[:i], c.views[i+1:]...)

			return nil
		}
	}

	return fmt.Errorf("view %d not found", identifier)
}
