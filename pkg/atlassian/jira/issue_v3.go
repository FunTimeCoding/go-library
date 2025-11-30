package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/web/parameter"
)

func (c *Client) IssueV3(key string) {
	status, response := c.basic.Get(
		c.basic.Base().Copy().Base(
			constant.Base,
		).Path("%s/%s", constant.Issue, key).Set(
			parameter.Fields,
			constant.AllFields,
		).String(),
	)
	fmt.Printf("Response: %d %s", status, response)
}
