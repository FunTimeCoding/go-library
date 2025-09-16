package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) IssueTypeFields(t *jira.MetaIssueType) map[string]string {
	result, e := t.GetAllFields()
	errors.PanicOnError(e)

	return result
}
