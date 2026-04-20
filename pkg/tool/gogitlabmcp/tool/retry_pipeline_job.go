package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type retryPipelineJobArguments struct {
	Project string `json:"project"`
	Job     int64  `json:"job"`
}

func (t *Tool) RetryPipelineJob(
	_ context.Context,
	_ mcp.CallToolRequest,
	a retryPipelineJobArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Job == 0 {
		return response.Fail("job is required")
	}

	v, _, e := t.client.Jobs.RetryJob(a.Project, a.Job)

	if e != nil {
		return response.Fail("retry pipeline job: %v", e)
	}

	return response.SuccessAny(v)
}
