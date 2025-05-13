package chromium

import (
	"github.com/funtimecoding/go-library/pkg/chromium/tab"
	"strings"
)

func (c *Client) TabByHost(s string) *tab.Tab {
	for _, t := range c.Tabs() {
		if strings.Contains(t.Url, s) {
			return t
		}
	}

	return nil
}
