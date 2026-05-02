package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetProjectVariable(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetProjectVariable,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Key == "" {
		return response.Fail("key is required")
	}

	v, _, e := s.client.ProjectVariables.GetVariable(
		a.Project,
		a.Key,
		nil,
	)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(v)
}
