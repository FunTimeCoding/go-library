package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool) RetryPipeline(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.RetryPipeline,
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
