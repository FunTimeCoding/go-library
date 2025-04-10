package basic_client

import (
	"fmt"
	"time"
)

func (c *Client) Series(series string) string {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)
	result := c.Get(
		fmt.Sprintf(
			"/series?start=%d&end=%d&match[]=%s",
			oneWeekAgo.Unix(),
			now.Unix(),
			series,
		),
	)
	fmt.Printf("Series: %s\n", result)

	return result
}
