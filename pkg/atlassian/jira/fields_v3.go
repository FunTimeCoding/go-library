package jira

import "fmt"

func (c *Client) FieldsV3() {
	status, response := c.basic.GetPath("/rest/api/3/field")
	// Does not contain more fields than the V2 API
	fmt.Printf("Basic response: %d %s", status, response)
}
