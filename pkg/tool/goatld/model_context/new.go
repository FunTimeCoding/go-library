package model_context

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	j *jira.Client,
	c *confluence.Client,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			"atlassian",
			constant.DefaultVersion,
			server.WithToolCapabilities(true),
		),
		jira:       j,
		confluence: c,
	}
	result.register()

	return result
}
