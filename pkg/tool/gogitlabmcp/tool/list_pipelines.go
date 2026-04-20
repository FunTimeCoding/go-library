package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type listPipelinesArguments struct {
	Project   string `json:"project"`
	Reference string `json:"reference"`
	Status    string `json:"status"`
}

func (t *Tool) ListPipelines(
	_ context.Context,
	_ mcp.CallToolRequest,
	a listPipelinesArguments,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	o := &gitlab.ListProjectPipelinesOptions{
		ListOptions: gitlab.ListOptions{PerPage: 20},
	}

	if a.Reference != "" {
		o.Ref = &a.Reference
	}

	if a.Status != "" {
		s := gitlab.BuildStateValue(a.Status)
		o.Status = &s
	}

	v, _, e := t.client.Pipelines.ListProjectPipelines(a.Project, o)

	if e != nil {
		return response.Fail("list pipelines: %v", e)
	}

	return response.SuccessAny(v)
}
