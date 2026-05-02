package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetMergeRequest(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetMergeRequest,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.MergeRequest == 0 {
		return response.Fail("merge_request is required")
	}

	v, _, e := s.client.MergeRequests.GetMergeRequest(
		a.Project,
		a.MergeRequest,
		nil,
	)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(v)
}
