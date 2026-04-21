package option

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
)

type Atlassian struct {
	Jira       *jira.Client
	Confluence *confluence.Client
	Port       int
}
