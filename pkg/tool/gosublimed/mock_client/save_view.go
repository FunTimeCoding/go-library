package mock_client

import "fmt"

func (c *Client) SaveView(
	identifier int,
	path string,
) error {
	for i, v := range c.views {
		if v.Identifier == identifier {
			c.views[i].FilePath = path
			c.views[i].Dirty = false

			return nil
		}
	}

	return fmt.Errorf("view %d not found", identifier)
}
