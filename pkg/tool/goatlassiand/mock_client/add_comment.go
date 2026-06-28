package mock_client

import "fmt"

func (c *Client) AddComment(
	pageIdentifier string,
	body string,
) error {
	e, okay := c.pages[pageIdentifier]

	if !okay || e.deleted {
		return fmt.Errorf("page not found: %s", pageIdentifier)
	}

	return nil
}
