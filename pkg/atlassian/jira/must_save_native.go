package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustSaveNative(i *jira.Issue) *issue.Issue {
	result, e := c.SaveNative(i)
	errors.PanicOnError(e)

	return result
}
