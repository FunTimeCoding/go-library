package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/web/parameter"
)

func (c *Client) CustomerIssuesBasic() {
	status, response := c.basic.Get(
		c.basic.Base().Copy().Base(
			constant.ServiceDesk,
		).Path(constant.Request).SetInteger(
			parameter.Limit,
			10,
		).SetInteger(parameter.Start, 0).String(),
	)
	fmt.Printf("Basic response: %d %s", status, response)
}
