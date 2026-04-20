package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type cancelPipelineArguments struct {
	Project  string `json:"project"`
	Pipeline int64  `json:"pipeline"`
}

func (t *Tool) CancelPipeline(
	_ context.Context,
	_ mcp.CallToolRequest,
	a cancelPipelineArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Pipeline == 0 {
		return response.Fail("pipeline is required")
	}

	v, _, e := t.client.Pipelines.CancelPipelineBuild(
		a.Project,
		a.Pipeline,
	)

	if e != nil {
		return response.Fail("cancel pipeline: %v", e)
	}

	return response.SuccessAny(v)
}
