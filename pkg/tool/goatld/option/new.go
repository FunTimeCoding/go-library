package option

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
)

func New() *Atlassian {
	return &Atlassian{
		Jira:       jira.NewEnvironment(),
		Confluence: confluence.NewEnvironment(),
	}
}
