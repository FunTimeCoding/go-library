package jira

import "fmt"

func (c *Client) CustomerIssuesBasic() {
	status, response := c.basic.Request(
		"/rest/servicedeskapi/request?limit=10&start=0",
	)
	fmt.Printf("Basic response: %d %s", status, response)
}
