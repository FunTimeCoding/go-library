package opsgenie

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/alert_enricher"

func (c *Client) Set(e *alert_enricher.Enricher) *Client {
	c.enricher = e

	return c
}
