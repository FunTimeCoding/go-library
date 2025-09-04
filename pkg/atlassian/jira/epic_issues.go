package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
)

func (c *Client) EpicIssues(name string) []*issue.Issue {
	return c.SearchFull("%s = %s", constant.ParentEpic, name)
}
