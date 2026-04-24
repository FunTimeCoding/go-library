package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (t *Tool) ListMergeRequests(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListMergeRequests,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	o := &gitlab.ListProjectMergeRequestsOptions{
		ListOptions: gitlab.ListOptions{PerPage: 20},
	}

	if a.State != "" {
		o.State = &a.State
	}

	v, _, e := t.client.MergeRequests.ListProjectMergeRequests(
		a.Project,
		o,
	)

	if e != nil {
		return response.Fail("list merge requests: %v", e)
	}

	return response.SuccessAny(v)
}
