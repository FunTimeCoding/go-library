package prometheus

import "time"

func (c *Client) QueryFloat(
	q string,
	t time.Time,
) float64 {
	result := c.QueryFloats(q, t)

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
