package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) EditMeta(key string) *jira.EditMetaInfo {
	result, _, e := c.client.Issue.GetEditMeta(&jira.Issue{Key: key})
	errors.PanicOnError(e)

	return result
}
