package jira

import "fmt"

func (c *Client) IssueV3(key string) {
	status, response := c.basic.Get(
		c.basic.Base().Copy().Base(
			"/rest/api/3",
		).Path("/issue/%s", key).Set(
			"fields",
			"*all",
		).String(),
	)
	fmt.Printf("Response: %d %s", status, response)
}
