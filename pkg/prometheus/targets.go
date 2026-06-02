package prometheus

import "github.com/prometheus/client_golang/api/prometheus/v1"

func (c *Client) Targets() (v1.TargetsResult, error) {
	return c.client.Targets(c.context)
}
