package tool

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type getCommitDiffArguments struct {
	Project string `json:"project"`
	Sha     string `json:"sha"`
}

func (t *Tool) GetCommitDiff(
	_ context.Context,
	_ mcp.CallToolRequest,
	a getCommitDiffArguments,
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
