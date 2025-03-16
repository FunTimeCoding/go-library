package loki

import "time"

func (c *Client) Labels(
	start time.Time,
	end time.Time,
) []string {

	return c.basic.Labels(start, end)
}
