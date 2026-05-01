package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustFillComments(v []*issue.Issue) {
	errors.PanicOnError(c.FillComments(v))
}
