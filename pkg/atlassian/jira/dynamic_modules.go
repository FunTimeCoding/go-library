package jira

import "fmt"

func (c *Client) DynamicModules() {
	status, response := c.basic.GetPath(
		"/rest/atlassian-connect/1/app/module/dynamic",
	)
	// 401 {"message":"The request is not from a Connect app."}
	fmt.Printf("Basic response: %d %s", status, response)
}
