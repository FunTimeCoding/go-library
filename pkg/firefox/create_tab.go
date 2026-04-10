package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/firefox/tab"
)

func (c *Client) CreateTab(l string) (tab.Tab, error) {
	r, e := c.send("create_tab", map[string]string{"url": l})

	if e != nil {
		return tab.Tab{}, fmt.Errorf("create tab: %w", e)
	}

	var result tab.Tab

	if e = json.Unmarshal(r.Result, &result); e != nil {
		return tab.Tab{}, fmt.Errorf("create tab: %w", e)
	}

	return result, nil
}
