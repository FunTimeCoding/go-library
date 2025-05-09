package jira

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"

func (c *Client) IssueOption() *option.Issue {
	if c.issueOption != nil {
		return c.issueOption
	}

	c.issueOption = option.New(
		c.locator,
		c.user,
		c.WatchedIssueKeys(),
		c.FieldMap(),
	)
	c.issueOption.Verbose = c.verbose

	return c.issueOption
}
