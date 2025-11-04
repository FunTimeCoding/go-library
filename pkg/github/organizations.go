package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v77/github"
)

func (c *Client) Organizations(user string) []*github.Organization {
	result, _, e := c.client.Organizations.List(c.context, user, nil)
	errors.PanicOnError(e)

	return result
}
