package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustEditMeta(key string) *jira.EditMetaInfo {
	result, e := c.EditMeta(key)
	errors.PanicOnError(e)

	return result
}
