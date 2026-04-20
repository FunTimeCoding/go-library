package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

type getPipelineJobArguments struct {
	Project string `json:"project"`
	Job     int64  `json:"job"`
}

func (t *Tool) GetPipelineJob(
	_ context.Context,
	_ mcp.CallToolRequest,
	a getPipelineJobArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Job == 0 {
		return response.Fail("job is required")
	}

	v, _, e := t.client.Jobs.GetJob(a.Project, a.Job)

	if e != nil {
		return response.Fail("get pipeline job: %v", e)
	}

	return response.SuccessAny(v)
}
