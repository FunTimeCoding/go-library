package jira

import "fmt"

func (c *Client) FieldsV3() error {
	status, body, e := c.basic.GetPath("/rest/api/3/field")

	if e != nil {
		return e
	}

	// Does not contain more fields than the V2 API
	fmt.Printf("Basic response: %d %s", status, body)

	return nil
}
