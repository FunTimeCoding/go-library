package model_context

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	j *jira.Client,
	c *confluence.Client,
	r face.Reporter,
	version string,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Identity.Name(),
			version,
			server.WithToolCapabilities(true),
			server.WithInstructions(constant.Identity.Instructions()),
		),
		jira:       j,
		confluence: c,
		reporter:   r,
	}
	result.register()

	return result
}
