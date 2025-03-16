package basic_client

import (
	"fmt"
	"time"
)

func (c *Client) QueryRange(query string) string {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)
	result := c.Get(
		fmt.Sprintf(
			"/query_range?start=%d&end=%d&query=%s",
			oneWeekAgo.Unix(),
			now.Unix(),
			query,
		),
	)
	fmt.Printf("Query range: %s\n", result)

	return result
}
