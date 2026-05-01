package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/customer"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustCustomerIssues(all bool) []*customer.Issue {
	result, e := c.CustomerIssues(all)
	errors.PanicOnError(e)

	return result
}
