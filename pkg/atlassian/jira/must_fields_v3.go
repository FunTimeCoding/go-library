package jira

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustFieldsV3() {
	errors.PanicOnError(c.FieldsV3())
}
