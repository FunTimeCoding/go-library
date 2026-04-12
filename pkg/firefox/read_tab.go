package firefox

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/firefox/tab"
)

func (c *Client) ReadTab(
	identifier int,
	raw bool,
) (tab.Content, error) {
	r, e := c.send(
		"read_tab", map[string]any{
			"tab_id": identifier,
			"raw":    raw,
		},
	)

	if e != nil {
		return tab.Content{}, fmt.Errorf("read tab: %w", e)
	}

	var result tab.Content

	if e = json.Unmarshal(r.Result, &result); e != nil {
		return tab.Content{}, fmt.Errorf("read tab: %w", e)
	}

	return result, nil
}
