package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/firefox/tab"
)

func (c *Client) Tabs() ([]tab.Tab, error) {
	r, e := c.send("list_tabs", nil)

	if e != nil {
		return nil, fmt.Errorf("list tabs: %w", e)
	}

	var result []tab.Tab

	if e = json.Unmarshal(r.Result, &result); e != nil {
		return nil, fmt.Errorf("list tabs: %w", e)
	}

	return result, nil
}
