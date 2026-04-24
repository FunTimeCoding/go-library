package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (t *Tool) ListCommits(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListCommits,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	o := &gitlab.ListCommitsOptions{
		ListOptions: gitlab.ListOptions{PerPage: 20},
	}

	if a.Reference != "" {
		o.RefName = &a.Reference
	}

	v, _, e := t.client.Commits.ListCommits(a.Project, o)

	if e != nil {
		return response.Fail("list commits: %v", e)
	}

	return response.SuccessAny(v)
}
