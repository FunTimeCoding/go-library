package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (t *Tool) CreatePipeline(
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

	v, _, e := t.client.Pipelines.CreatePipeline(
		a.Project,
		&gitlab.CreatePipelineOptions{Ref: &a.Reference},
	)

	if e != nil {
		return response.Fail("create pipeline: %v", e)
	}

	return response.SuccessAny(v)
}
