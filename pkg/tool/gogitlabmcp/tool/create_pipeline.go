package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type createPipelineArguments struct {
	Project   string `json:"project"`
	Reference string `json:"reference"`
}

func (t *Tool) CreatePipeline(
	_ context.Context,
	_ mcp.CallToolRequest,
	a createPipelineArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Reference == "" {
		return response.Fail("reference is required")
	}

	v, _, e := t.client.Pipelines.CreatePipeline(
		a.Project,
		&gitlab.CreatePipelineOptions{Ref: &a.Reference},
	)

	if e != nil {
		return response.Fail("create pipeline: %v", e)
	}

	return response.SuccessAny(v)
}
