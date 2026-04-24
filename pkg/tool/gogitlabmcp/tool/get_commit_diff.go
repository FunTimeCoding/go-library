package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (t *Tool) GetCommitDiff(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetCommitDiff,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Sha == "" {
		return response.Fail("sha is required")
	}

	v, _, e := t.client.Commits.GetCommitDiff(
		a.Project,
		a.Sha,
		&gitlab.GetCommitDiffOptions{
			ListOptions: gitlab.ListOptions{PerPage: 100},
		},
	)

	if e != nil {
		return response.Fail("get commit diff: %v", e)
	}

	return response.SuccessAny(v)
}
