package prometheus

import (
	"fmt"
	"net/url"
)

func (c *Client) GraphLink(expression string) string {
	return fmt.Sprintf(
		"https://%s/graph?g0.expr=%s",
		c.host,
		url.QueryEscape(expression),
	)
}
