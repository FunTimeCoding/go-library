package loki

import "time"

func (c *Client) LabelValues(
	start time.Time,
	end time.Time,
	label string,
) []string {
	return c.basic.LabelValues(start, end, label)
}
