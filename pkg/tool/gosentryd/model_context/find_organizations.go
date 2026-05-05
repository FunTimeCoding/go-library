package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) FindOrganizations(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.FindOrganizations,
) (*mcp.CallToolResult, error) {
	result, e := s.client.Organizations()

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(result)
}
