package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
)

func (c *Client) DynamicModules() {
	status, response := c.basic.GetPath(constant.Dynamic)
	// 401 {"message":"The request is not from a Connect app."}
	fmt.Printf("DynamicModule: %d %s\n", status, response)
}
