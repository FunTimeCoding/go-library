package prometheus

import "time"

func (c *Client) IntegerExists(
	q string,
	t time.Time,
) bool {
	return len(c.QueryIntegers(q, t)) > 0
}
