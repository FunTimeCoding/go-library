package mock_client

import "github.com/funtimecoding/go-library/pkg/firefox/tab"

func (c *Client) CreateTab(l string) (*tab.Tab, error) {
	t := tab.New()
	t.Identifier = c.nextID
	t.Locator = l
	t.Status = "complete"
	c.nextID++
	c.tabs = append(c.tabs, t)

	return t, nil
}
