package jira

import "fmt"

func (c *Client) Addons() {
	status, response := c.basic.Request(
		"/rest/atlassian-connect/1/addons",
	)
	// 403 {"message":"Client must be authenticated as a system administrator to access this resource.","status-code":403}
	fmt.Printf("Basic response: %d %s", status, response)
}
