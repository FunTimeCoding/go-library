package mock_client

import "github.com/funtimecoding/go-library/pkg/firefox/tab"

func (c *Client) AddTab(t *tab.Tab) {
	c.tabs = append(c.tabs, t)
}
