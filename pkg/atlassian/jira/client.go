package jira

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
)

type Client struct {
	context context.Context
	client  *jira.Client
	locator string
	user    string

	fieldMap    *field_map.Map
	issueOption *option.Issue
}
