package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) CreateMeta(key string) *jira.CreateMetaInfo {
	result, _, e := c.client.Issue.GetCreateMeta(key)
	errors.PanicOnError(e)

	return result
}
