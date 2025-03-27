package confluence

import (
	"fmt"
	"net/url"
)

func (c *Client) SearchBasic(query string) string {
	return c.basic.Get(
		fmt.Sprintf("/content/search?cql=%s", url.QueryEscape(query)),
	)
}
