package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
)

func (c *Client) Addons() {
	status, response := c.basic.GetPath(constant.Addon)
	// 403 {"message":"Client must be authenticated as a system administrator to access this resource.","status-code":403}
	fmt.Printf("Addon: %d %s\n", status, response)
}
