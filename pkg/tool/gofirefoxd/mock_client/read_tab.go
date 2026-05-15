package mock_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/firefox/tab"
)

func (c *Client) ReadTab(
	identifier int,
	_ bool,
) (tab.Content, error) {
	for _, t := range c.tabs {
		if t.Identifier == identifier {
			return tab.Content{
				Identifier: t.Identifier,
				Locator:    t.Locator,
				Title:      t.Title,
			}, nil
		}
	}

	return tab.Content{}, fmt.Errorf("tab %d not found", identifier)
}
