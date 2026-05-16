package mock_client

import "github.com/funtimecoding/go-library/pkg/firefox/tab"

func (c *Client) Tabs() ([]*tab.Tab, error) {
	return c.tabs, nil
}
