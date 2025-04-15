package jira

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/ctreminiom/go-atlassian/jira/sm"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/basic_client"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
)

type Client struct {
	context context.Context
	client  *jira.Client
	locator string
	user    string

	basic   *basic_client.Client
	service *sm.Client

	fieldMap    *field_map.Map
	issueOption *option.Issue
}
