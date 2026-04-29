package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v85/github"
)

func (c *Client) MustOrganizations(user string) []*github.Organization {
	result, e := c.Organizations(user)
	errors.PanicOnError(e)

	return result
}
