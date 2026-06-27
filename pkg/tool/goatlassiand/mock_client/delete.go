package mock_client

import "fmt"

func (c *Client) Delete(pageIdentifier string) error {
	e, okay := c.pages[pageIdentifier]

	if !okay || e.deleted || e.page == nil || e.page.Status != "current" {
		return fmt.Errorf("page not found: %s", pageIdentifier)
	}

	e.deleted = true

	return nil
}
