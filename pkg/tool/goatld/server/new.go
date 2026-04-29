package server

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
)

func New(
	l *jira.Client,
	c *confluence.Client,
) *Server {
	return &Server{jira: l, confluence: c}
}
