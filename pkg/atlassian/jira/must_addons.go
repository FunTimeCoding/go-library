package jira

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustAddons() {
	errors.PanicOnError(c.Addons())
}
