package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustIssueOption() *option.Issue {
	result, e := c.IssueOption()
	errors.PanicOnError(e)

	return result
}
