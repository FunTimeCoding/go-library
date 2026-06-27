package mock_client

import "fmt"

func (c *Client) DeleteDraft(pageIdentifier string) error {
	e, okay := c.pages[pageIdentifier]

	if !okay || e.deleted {
		return fmt.Errorf("page not found: %s", pageIdentifier)
	}

	if e.draft != nil {
		e.draft = nil

		return nil
	}

	if e.page != nil && e.page.Status == "draft" {
		e.deleted = true

		return nil
	}

	return fmt.Errorf("no draft to delete: %s", pageIdentifier)
}
