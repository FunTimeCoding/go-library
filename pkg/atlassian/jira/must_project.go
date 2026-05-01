package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustProject(key string) *jira.Project {
	result, e := c.Project(key)
	errors.PanicOnError(e)

	return result
}
