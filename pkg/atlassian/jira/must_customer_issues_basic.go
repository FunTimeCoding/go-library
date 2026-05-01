package jira

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustCustomerIssuesBasic() {
	errors.PanicOnError(c.CustomerIssuesBasic())
}
