package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) FindProjects(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.FindProjects,
) (*mcp.CallToolResult, error) {
	result, e := s.client.OrganizationProjects(s.organization)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(result)
}
