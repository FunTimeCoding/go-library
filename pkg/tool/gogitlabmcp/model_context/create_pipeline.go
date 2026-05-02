package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) CreatePipeline(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreatePipeline,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Reference == "" {
		return response.Fail("reference is required")
	}

	v, _, e := s.client.Pipelines.CreatePipeline(
		a.Project,
		&gitlab.CreatePipelineOptions{Ref: &a.Reference},
	)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(v)
}
