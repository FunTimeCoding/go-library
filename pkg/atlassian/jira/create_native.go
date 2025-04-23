package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
)

func (c *Client) CreateNative(i *jira.Issue) *issue.Issue {
	result, r, e := c.client.Issue.CreateWithContext(c.context, i)

	if e != nil && r != nil {
		fmt.Printf("Body: %s\n", string(system.ReadAll(r.Body)))
	}

	errors.PanicOnError(e)

	return c.enrichOne(issue.New(result, c.IssueOption()))
}
