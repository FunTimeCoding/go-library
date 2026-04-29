package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v85/github"
)

func (c *Client) MustProjectIssues(
	owner string,
	name string,
) []*github.Issue {
	result, e := c.ProjectIssues(owner, name)
	errors.PanicOnError(e)

	return result
}
