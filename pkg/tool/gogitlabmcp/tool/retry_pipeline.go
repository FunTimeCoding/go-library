package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type retryPipelineArguments struct {
	Project  string `json:"project"`
	Pipeline int64  `json:"pipeline"`
}

func (t *Tool) RetryPipeline(
	_ context.Context,
	_ mcp.CallToolRequest,
	a retryPipelineArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Pipeline == 0 {
		return response.Fail("pipeline is required")
	}

	v, _, e := t.client.Pipelines.RetryPipelineBuild(
		a.Project,
		a.Pipeline,
	)

	if e != nil {
		return response.Fail("retry pipeline: %v", e)
	}

	return response.SuccessAny(v)
}
