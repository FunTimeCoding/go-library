package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
)

func (t *Tool) GetPipelineJobOutput(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetPipelineJobOutput,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Job == 0 {
		return response.Fail("job is required")
	}

	r, _, e := t.client.Jobs.GetTraceFile(a.Project, a.Job)

	if e != nil {
		return response.Fail("get pipeline job output: %v", e)
	}

	b, e := io.ReadAll(r)

	if e != nil {
		return response.Fail("read job output: %v", e)
	}

	return mcp.NewToolResultText(string(b)), nil
}
