package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) UpdateProjectVariable(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateProjectVariable,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Key == "" {
		return response.Fail("key is required")
	}

	if a.Value == "" {
		return response.Fail("value is required")
	}

	v, _, e := s.client.ProjectVariables.UpdateVariable(
		a.Project,
		a.Key,
		&gitlab.UpdateProjectVariableOptions{
			Value:     &a.Value,
			Protected: &a.Protected,
			Masked:    &a.Masked,
		},
	)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(v)
}
