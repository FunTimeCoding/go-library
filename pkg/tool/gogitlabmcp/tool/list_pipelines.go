package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (t *Tool) ListPipelines(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListPipelines,
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
