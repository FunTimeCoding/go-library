package prometheus

import (
	"fmt"
	"net/url"
)

func (c *Client) GraphLink(expression string) string {
	var host string

	if c.alternateGraphHost != "" {
		host = c.alternateGraphHost
	} else {
		host = c.host
	}

	return fmt.Sprintf(
		"https://%s/graph?g0.expr=%s",
		host,
		url.QueryEscape(expression),
	)
}
