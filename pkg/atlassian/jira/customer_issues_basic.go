package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
)

func (c *Client) CustomerIssuesBasic() {
	status, response := c.basic.Get(
		c.basic.Base().Copy().Base(
			constant.ServiceDesk,
		).Path(constant.Request).SetInteger(
			constant.LimitKey,
			10,
		).SetInteger(constant.StartKey, 0).String(),
	)
	fmt.Printf("Basic response: %d %s", status, response)
}
