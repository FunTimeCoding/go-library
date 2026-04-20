package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type listProjectsArguments struct {
	Search string `json:"search"`
}

func (t *Tool) ListProjects(
	_ context.Context,
	_ mcp.CallToolRequest,
	a listProjectsArguments,
) (*mcp.CallToolResult, error) {
	o := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{PerPage: 20},
	}

	if a.Search != "" {
		o.Search = &a.Search
	}

	v, _, e := t.client.Projects.ListProjects(o)

	if e != nil {
		return response.Fail("list projects: %v", e)
	}

	return response.SuccessAny(v)
}
