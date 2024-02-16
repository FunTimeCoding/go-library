package github

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github/issue"
)

func (c *Client) OpenUserIssues() []*issue.Issue {
	return c.Search(
		fmt.Sprintf("is:open is:issue author:%s", *c.User().Login),
	)
}
