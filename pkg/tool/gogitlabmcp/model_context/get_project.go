package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetProject(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetProject,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	v, _, e := s.client.Projects.GetProject(a.Project, nil)

	if e != nil {
		return s.captureFail(e, "get project")
	}

	return response.SuccessAny(v)
}
