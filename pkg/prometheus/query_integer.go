package prometheus

import "time"

func (c *Client) QueryInteger(
	q string,
	t time.Time,
) int {
	result := c.QueryIntegers(q, t)

	if len(result) == 0 {
		return 0
	}

	if len(result) > 1 {
		panic("more than one result")
	}

	for _, v := range result {
		return v
	}

	return 0
}
