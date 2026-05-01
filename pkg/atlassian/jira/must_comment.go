package jira

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustComment(
	key string,
	body string,
) {
	errors.PanicOnError(c.Comment(key, body))
}
