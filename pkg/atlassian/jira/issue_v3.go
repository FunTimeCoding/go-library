package jira

import "fmt"

func (c *Client) IssueV3(key string) {
	status, response := c.basic.Get(
		fmt.Sprintf("/rest/api/3/issue/%s?fields=*all", key),
	)
	fmt.Printf("Response: %d %s", status, response)
}
