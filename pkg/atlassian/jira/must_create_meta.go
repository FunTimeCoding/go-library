package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustCreateMeta(key string) *jira.CreateMetaInfo {
	result, e := c.CreateMeta(key)
	errors.PanicOnError(e)

	return result
}
