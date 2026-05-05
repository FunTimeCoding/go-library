package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) ListPipelineJobs(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListPipelineJobs,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Pipeline == 0 {
		return response.Fail("pipeline is required")
	}

	v, _, e := s.client.Jobs.ListPipelineJobs(
		a.Project,
		a.Pipeline,
		&gitlab.ListJobsOptions{
			ListOptions: gitlab.ListOptions{PerPage: 100},
		},
	)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(v)
}
