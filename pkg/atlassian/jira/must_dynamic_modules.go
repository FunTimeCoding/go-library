package jira

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDynamicModules() {
	errors.PanicOnError(c.DynamicModules())
}
