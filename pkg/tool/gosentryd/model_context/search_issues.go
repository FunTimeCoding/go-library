package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SearchIssues(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchIssues,
) (*mcp.CallToolResult, error) {
	result, e := s.client.SearchIssues(
		s.organization,
		a.Query,
		a.Project,
		a.Limit,
		a.Cursor,
	)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(result)
}
