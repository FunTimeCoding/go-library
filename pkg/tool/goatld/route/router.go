package route

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
)

type Router struct {
	jira       *jira.Client
	confluence *confluence.Client
}
