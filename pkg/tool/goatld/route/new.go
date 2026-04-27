package route

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
)

func New(
	l *jira.Client,
	c *confluence.Client,
) *Router {
	return &Router{jira: l, confluence: c}
}
