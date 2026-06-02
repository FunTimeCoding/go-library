package prometheus

import "github.com/prometheus/client_golang/api/prometheus/v1"

func (c *Client) Metadata(metric string) (map[string][]v1.Metadata, error) {
	return c.client.Metadata(c.context, metric, "")
}
