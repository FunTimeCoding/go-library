package basic_client

import (
	"fmt"
	"net/url"
	"time"
)

func (c *Client) Query(query string) string {
	result := c.Get(
		fmt.Sprintf(
			"/query?time=%d&query=%s",
			time.Now().Unix(),
			url.QueryEscape(query),
		),
	)

	return result
}
