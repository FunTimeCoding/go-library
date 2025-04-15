package jira

import "fmt"

func (c *Client) IssueV3(key string) {
	status, response := c.basic.Request(
		fmt.Sprintf("/rest/api/3/issue/%s?fields=*all", key),
	)
	fmt.Printf("Basic response: %d %s", status, response)
}
