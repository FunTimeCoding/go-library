package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Issue(key string) *issue.Issue {
	result, r, e := c.client.Issue.Get(key, &jira.GetQueryOptions{})

	if e != nil && r != nil {
		if r.StatusCode >= 400 {
			fmt.Println("Status:", r.Status)
			fmt.Println("Response:", r)
			fmt.Println("Error:", e.Error())
			web.PrintHeader(r.Header)
		}
	}

	errors.PanicOnError(e)

	return issue.New(result, c.IssueOption())
}
