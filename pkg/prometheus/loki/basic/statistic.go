package basic

import (
	"fmt"
	"time"
)

func (c *Client) Statistic(query string) string {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)
	result := c.Get(
		fmt.Sprintf(
			"/index/stats?start=%d&end=%d&index=%s",
			oneWeekAgo.Unix(),
			now.Unix(),
			query,
		),
	)
	fmt.Printf("Index stats: %s\n", result)

	return result
}
