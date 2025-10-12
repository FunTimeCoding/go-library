package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
)

func (c *Client) IssueV3(key string) {
	status, response := c.basic.Get(
		c.basic.Base().Copy().Base(
			constant.Base,
		).Path("%s/%s", constant.Issue, key).Set(
			constant.FieldsKey,
			constant.AllFields,
		).String(),
	)
	fmt.Printf("Response: %d %s", status, response)
}
