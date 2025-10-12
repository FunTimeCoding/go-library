package jira

import "fmt"

func (c *Client) CustomerIssuesBasic() {
	status, response := c.basic.Get(
		c.basic.Base().Copy().Base(
			"/rest/servicedeskapi",
		).Path("/request").Set(
			"limit",
			"10",
		).Set("start", "0").String(),
	)
	fmt.Printf("Basic response: %d %s", status, response)
}
