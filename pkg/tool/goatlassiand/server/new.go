package server

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/face"
)

func New(
	l *jira.Client,
	c *confluence.Client,
	r face.Reporter,
) *Server {
	return &Server{
		jira:       l,
		confluence: c,
		reporter:   r,
	}
}
