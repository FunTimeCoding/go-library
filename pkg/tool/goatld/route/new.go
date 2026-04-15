package route

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
)

func New(
	j *jira.Client,
	c *confluence.Client,
) *Router {
	return &Router{jira: j, confluence: c}
}
